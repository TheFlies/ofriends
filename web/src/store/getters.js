const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.auth.token,
  // avatar: state => state.user.avatar,
  name: state => state.auth.name,
  roles: state => state.auth.roles,
  info: state => state.auth.info
}
export default getters

