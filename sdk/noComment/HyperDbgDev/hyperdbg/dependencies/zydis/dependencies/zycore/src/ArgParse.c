#include <Zycore/ArgParse.h>
#include <Zycore/LibC.h>
#ifndef ZYAN_NO_LIBC
ZyanStatus ZyanArgParse(const ZyanArgParseConfig *cfg, ZyanVector *parsed,
                        const char **error_token) {
  return ZyanArgParseEx(cfg, parsed, error_token, ZyanAllocatorDefault());
}

#endif
ZyanStatus ZyanArgParseEx(const ZyanArgParseConfig *cfg, ZyanVector *parsed,
                          const char **error_token, ZyanAllocator *allocator) {
#define ZYAN_ERR_TOK(tok)                                                      \
  if (error_token) {                                                           \
    *error_token = tok;                                                        \
  }
  ZYAN_ASSERT(cfg);
  ZYAN_ASSERT(parsed);
  if (cfg->min_unnamed_args > cfg->max_unnamed_args) {
    return ZYAN_STATUS_INVALID_ARGUMENT;
  }
  for (const ZyanArgParseDefinition *def = cfg->args; def && def->name; ++def) {
    if (!def->name) {
      return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize arg_len = ZYAN_STRLEN(def->name);
    if (arg_len < 2 || def->name[0] != '-') {
      return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (def->name[1] != '-' && arg_len != 2) {
      return ZYAN_STATUS_INVALID_ARGUMENT;
    }
  }
  ZYAN_CHECK(ZyanVectorInitEx(
      parsed, sizeof(ZyanArgParseArg), cfg->argc, ZYAN_NULL, allocator,
      ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR, ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD));
  ZyanStatus err;
  ZyanBool accept_dash_args = ZYAN_TRUE;
  ZyanUSize num_unnamed_args = 0;
  for (ZyanUSize i = 1; i < cfg->argc; ++i) {
    const char *cur_arg = cfg->argv[i];
    ZyanUSize arg_len = ZYAN_STRLEN(cfg->argv[i]);
    if (accept_dash_args && arg_len >= 2 &&
        ZYAN_MEMCMP(cur_arg, "--", 2) == 0) {
      if (arg_len == 2) {
        accept_dash_args = ZYAN_FALSE;
      }
      else {
        ZyanArgParseArg *parsed_arg;
        ZYAN_CHECK(ZyanVectorEmplace(parsed, (void **)&parsed_arg, ZYAN_NULL));
        ZYAN_MEMSET(parsed_arg, 0, sizeof(*parsed_arg));
        for (const ZyanArgParseDefinition *def = cfg->args; def && def->name;
             ++def) {
          if (ZYAN_STRCMP(def->name, cur_arg) == 0) {
            parsed_arg->def = def;
            break;
          }
        }
        if (!parsed_arg->def) {
          err = ZYAN_STATUS_ARG_NOT_UNDERSTOOD;
          ZYAN_ERR_TOK(cur_arg);
          goto failure;
        }
        if (!parsed_arg->def->boolean) {
          if (i == cfg->argc - 1) {
            err = ZYAN_STATUS_ARG_MISSES_VALUE;
            ZYAN_ERR_TOK(cur_arg);
            goto failure;
          }
          parsed_arg->has_value = ZYAN_TRUE;
          ZYAN_CHECK(
              ZyanStringViewInsideBuffer(&parsed_arg->value, cfg->argv[++i]));
        }
      }
      continue;
    }
    if (accept_dash_args && arg_len > 1 && cur_arg[0] == '-') {
      for (const char *read_ptr = cur_arg + 1; *read_ptr; ++read_ptr) {
        ZyanArgParseArg *parsed_arg;
        ZYAN_CHECK(ZyanVectorEmplace(parsed, (void **)&parsed_arg, ZYAN_NULL));
        ZYAN_MEMSET(parsed_arg, 0, sizeof(*parsed_arg));
        for (const ZyanArgParseDefinition *def = cfg->args; def && def->name;
             ++def) {
          if (ZYAN_STRLEN(def->name) == 2 && def->name[0] == '-' &&
              def->name[1] == *read_ptr) {
            parsed_arg->def = def;
            break;
          }
        }
        if (!parsed_arg->def) {
          err = ZYAN_STATUS_ARG_NOT_UNDERSTOOD;
          ZYAN_ERR_TOK(cur_arg);
          goto failure;
        }
        if (!parsed_arg->def->boolean) {
          if (read_ptr[1]) {
            parsed_arg->has_value = ZYAN_TRUE;
            ZYAN_CHECK(
                ZyanStringViewInsideBuffer(&parsed_arg->value, read_ptr + 1));
          }
          else {
            if (i == cfg->argc - 1) {
              err = ZYAN_STATUS_ARG_MISSES_VALUE;
              ZYAN_ERR_TOK(cur_arg)
              goto failure;
            }
            parsed_arg->has_value = ZYAN_TRUE;
            ZYAN_CHECK(
                ZyanStringViewInsideBuffer(&parsed_arg->value, cfg->argv[++i]));
          }
          goto continue_main_loop;
        }
      }
    }
    ++num_unnamed_args;
    if (num_unnamed_args > cfg->max_unnamed_args) {
      err = ZYAN_STATUS_TOO_MANY_ARGS;
      ZYAN_ERR_TOK(cur_arg);
      goto failure;
    }
    ZyanArgParseArg *parsed_arg;
    ZYAN_CHECK(ZyanVectorEmplace(parsed, (void **)&parsed_arg, ZYAN_NULL));
    ZYAN_MEMSET(parsed_arg, 0, sizeof(*parsed_arg));
    parsed_arg->has_value = ZYAN_TRUE;
    ZYAN_CHECK(ZyanStringViewInsideBuffer(&parsed_arg->value, cur_arg));
  continue_main_loop:;
  }
  if (num_unnamed_args < cfg->min_unnamed_args) {
    err = ZYAN_STATUS_TOO_FEW_ARGS;
    goto failure;
  }
  ZyanUSize num_parsed_args;
  ZYAN_CHECK(ZyanVectorGetSize(parsed, &num_parsed_args));
  for (const ZyanArgParseDefinition *def = cfg->args; def && def->name; ++def) {
    if (!def->required)
      continue;
    ZyanBool arg_found = ZYAN_FALSE;
    for (ZyanUSize i = 0; i < num_parsed_args; ++i) {
      const ZyanArgParseArg *arg = ZYAN_NULL;
      ZYAN_CHECK(ZyanVectorGetPointer(parsed, i, (const void **)&arg));
      if (!arg->def)
        continue;
      if (arg->def == def) {
        arg_found = ZYAN_TRUE;
        break;
      }
    }
    if (!arg_found) {
      err = ZYAN_STATUS_REQUIRED_ARG_MISSING;
      ZYAN_ERR_TOK(def->name);
      goto failure;
    }
  }
  ZYAN_ERR_TOK(ZYAN_NULL);
  return ZYAN_STATUS_SUCCESS;
failure:
  ZYAN_CHECK(ZyanVectorDestroy(parsed));
  return err;
#undef ZYAN_ERR_TOK
}

