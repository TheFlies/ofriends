const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  // token: state => state.user.token,
  // avatar: state => state.user.avatar,
  name: state => state.auth.name,
  roles: state => state.auth.roles
}
export default getters
