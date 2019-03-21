export function isvalidUsername(str) {
  return str.length > 4
  // const valid_map = ['admin', 'editor']
  // return valid_map.indexOf(str.trim()) >= 0
}

export function isExternal(path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}
