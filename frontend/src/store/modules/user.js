const user = {
  state: {
    token: '',
    userList: []
  },
  mutations: {
    SET_TOKEN(state, data) {
      state.token = data.token
    }
  },
  actions: {
    modifyToken({ commit }, token) {
      commit('SET_TOKEN', { token: token })
    }
  }
}

export default user
