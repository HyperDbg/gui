sudo pacman -S firefox
sudo pacman -R xfce4-screensaver
mount EWDK_ni_release_22621_220506-1250.iso -o loop /mnt
sudo mount EWDK.iso -o loop /mnt
sudo  mount -t iso9660 -o loop EWDK.iso /mnt

sudo pacman -S xarchiver