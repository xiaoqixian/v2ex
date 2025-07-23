// Date:   Sun Jul 13 16:31:53 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

const defaultAvatar = new URL("@/assets/default_avatar.png", import.meta.url).href

export function onAvatarError(event) {
  event.target.src = defaultAvatar
}
