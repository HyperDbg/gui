#include <Zydis/Internal/String.h>
#include <Zydis/FormatterBuffer.h>
ZyanStatus
ZydisFormatterTokenGetValue(const ZydisFormatterToken * token,
                            ZydisTokenType *            type,
                            ZyanConstCharPointer *      value) {
    if (!token || !type || !value) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *type  = token->type;
    *value = (ZyanConstCharPointer)((ZyanU8 *)token + sizeof(ZydisFormatterToken));
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterTokenNext(ZydisFormatterTokenConst ** token) {
    if (!token || !*token) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanU8 next = (*token)->next;
    if (!next) {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    *token = (ZydisFormatterTokenConst *)((ZyanU8 *)*token + sizeof(ZydisFormatterToken) + next);
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterBufferGetToken(const ZydisFormatterBuffer * buffer,
                             ZydisFormatterTokenConst **  token) {
    if (!buffer || !token) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *token = ((ZydisFormatterTokenConst *)buffer->string.vector.data - 1);
    if ((*token)->type == ZYDIS_TOKEN_INVALID) {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterBufferGetString(ZydisFormatterBuffer * buffer, ZyanString ** string) {
    if (!buffer || !string) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (buffer->is_token_list &&
        ((ZydisFormatterTokenConst *)buffer->string.vector.data - 1)->type == ZYDIS_TOKEN_INVALID) {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
    ZYAN_ASSERT(buffer->string.vector.data);
    ZYAN_ASSERT(buffer->string.vector.size);
    *string = &buffer->string;
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterBufferAppend(ZydisFormatterBuffer * buffer, ZydisTokenType type) {
    if (!buffer) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (!buffer->is_token_list) {
        return ZYAN_STATUS_SUCCESS;
    }
    const ZyanUSize len = buffer->string.vector.size;
    ZYAN_ASSERT((len > 0) && (len < 256));
    if (buffer->capacity <= len + sizeof(ZydisFormatterToken)) {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZydisFormatterToken * const last = (ZydisFormatterToken *)buffer->string.vector.data - 1;
    last->next                       = (ZyanU8)len;
    const ZyanUSize delta            = len + sizeof(ZydisFormatterToken);
    buffer->capacity -= delta;
    buffer->string.vector.data          = (ZyanU8 *)buffer->string.vector.data + delta;
    buffer->string.vector.size          = 1;
    buffer->string.vector.capacity      = ZYAN_MIN(buffer->capacity, 255);
    *(char *)buffer->string.vector.data = '\0';
    ZydisFormatterToken * const token   = (ZydisFormatterToken *)buffer->string.vector.data - 1;
    token->type                         = type;
    token->next                         = 0;
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterBufferRemember(const ZydisFormatterBuffer * buffer, ZyanUPointer * state) {
    if (!buffer || !state) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (buffer->is_token_list) {
        *state = (ZyanUPointer)buffer->string.vector.data;
    } else {
        *state = (ZyanUPointer)buffer->string.vector.size;
    }
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisFormatterBufferRestore(ZydisFormatterBuffer * buffer, ZyanUPointer state) {
    if (!buffer) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (buffer->is_token_list) {
        const ZyanUSize delta = (ZyanUPointer)buffer->string.vector.data - state;
        buffer->capacity += delta;
        buffer->string.vector.data          = (void *)state;
        buffer->string.vector.size          = 1; // TODO: Restore size?
        buffer->string.vector.capacity      = ZYAN_MIN(buffer->capacity, 255);
        *(char *)buffer->string.vector.data = '\0';
    } else {
        buffer->string.vector.size = (ZyanUSize)state;
        ZYDIS_STRING_NULLTERMINATE(&buffer->string);
    }
    return ZYAN_STATUS_SUCCESS;
}
