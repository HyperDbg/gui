mkdir -p .devcontainer
cd .devcontainer
cat >devcontainer.json<<-'EOFJSON'
// For format details, see https://aka.ms/devcontainer.json.
{
    "name": "Manjaro",
    "dockerFile": "Dockerfile",
    "runArgs": [
        "--cap-add=SYS_PTRACE",
        "--security-opt",
        "seccomp=unconfined"
    ],
    // "mounts": [
    //     "source=dind-var-lib-docker,target=/var/lib/docker,type=volume"
    // ],
    "mounts": [
        "source=/var/run/docker.sock,target=/var/run/docker.sock,type=bind"
    ],
    "overrideCommand": false,
    // Configure tool-specific properties.
    "customizations": {
        // Configure properties specific to VS Code.
        "vscode": {
            // Add the IDs of extensions you want installed when the container is created.
            "extensions": [
                // "MS-CEINTL.vscode-language-pack-zh-hans",
                "ms-azuretools.vscode-docker"
            ]
        }
    },
    // Use 'forwardPorts' to make a list of ports inside the container available locally.
    "forwardPorts": [
        5902
    ],
    // Use 'postCreateCommand' to run commands after the container is created.
    // "postCreateCommand": "docker --version",
    // Comment out to connect as root instead. More info: https://aka.ms/vscode-remote/containers/non-root.
    // "build": {
    //     "args": {
    //         "ENABLE_NONROOT_DOCKER": "false"
    //     }
    // },
    "remoteUser": "ddk"
}
EOFJSON

cat > Dockerfile<<-'EOFDKF'
# syntax=docker/dockerfile:1
#---------------------------
# FROM cake233/manjaro-zsh-amd64

FROM cake233/manjaro-xfce-amd64

# set username & group
ARG USERNAME=ddk
ARG GROUPNAME=ddk
# ARG USER_UID=1001
# ARG USER_GID=$USER_UID

# rm cn mirrorlist
RUN sed -e '/bfsu.edu.cn/d' \
    -e '/tuna.tsinghua.edu.cn/d' \
    -e '/opentuna.cn/d' \
    -i /etc/pacman.conf

# install dependencies
# live server: https://docs.microsoft.com/en-us/visualstudio/liveshare/reference/linux#install-linux-prerequisites
RUN pacman -Syu \
    --noconfirm \
    --needed \
    base \
    base-devel \
    git \
    lib32-gcc-libs \
    lib32-glibc \
    gcr \
    liburcu \
    openssl-1.0 \
    krb5 \
    icu \
    zlib \
    gnome-keyring \
    libsecret \
    desktop-file-utils \
    xorg-xprop \
    xdg-utils

# locale: Chinese Simplified (China)
ENV LANG=zh_CN.UTF-8

# add new user
RUN groupadd --force ${GROUPNAME} \
    && useradd --create-home --gid ${GROUPNAME} ${USERNAME} \
    && mkdir -p /etc/sudoers.d \
    && echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/ddk \
    && chmod 400 /etc/sudoers.d/ddk

WORKDIR ["/home/$USERNAME"]

# clean cache
RUN yes | pacman -Scc; \
    rm -rf /var/cache/pacman/pkg/* \
    /tmp/* \
    2>/dev/null

# command: sleep infinity
CMD [ "sleep", "inf" ]
EOFDKF
