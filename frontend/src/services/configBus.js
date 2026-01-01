const listeners = new Set()

export function registerReload(fn) {
  listeners.add(fn)
}

export function unregisterReload(fn) {
  listeners.delete(fn)
}

export function reloadAllTabs() {
  listeners.forEach(fn => fn())
}
