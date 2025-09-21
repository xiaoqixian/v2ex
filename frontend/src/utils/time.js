// Date:   Fri Jul 11 22:47:32 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

export function timeEval(seconds) {
  const now = Math.floor(Date.now() / 1000)
  const diff = now - seconds

  if (diff < 60) return '刚刚'
  if (diff < 3600) return `${Math.floor(diff / 60)} 分钟前`
  if (diff < 86400) return `${Math.floor(diff / 3600)} 小时前`
  if (diff < 2592000) return `${Math.floor(diff / 86400)} 天前`
  if (diff < 31536000) return `${Math.floor(diff / 2592000)} 个月前`
  return `${Math.floor(diff / 31536000)} 年前`
}

export function timeEvalTimeStr(timeStr) {
  return timeEval(Math.floor(new Date(timeStr).getTime() / 1000));
}
