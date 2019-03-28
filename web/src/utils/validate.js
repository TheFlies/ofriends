export function isvalidUsername(str) {
  return str && str.length > 4
  // const valid_map = ['admin', 'editor']
  // return valid_map.indexOf(str.trim()) >= 0
}

export function isvalidEmail(str) {
  var re = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/
  const email = String(str).toLowerCase()
  return re.test(email)// && email.endsWith('tma.com.vn')
}

export function isExternal(path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}
