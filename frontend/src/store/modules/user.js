const user = {
  state: {
    token: '',
    nickName: ''
  },
  mutations: {
    SET_TOKEN(state, data) {
      state.token = data.token
    },
    SET_NICKNAME(state, data) {
      state.nickName = data;
    }
  },
  actions: {
    modifyToken({ commit }, token) {
      commit('SET_TOKEN', { token: token })
    },
    modifyName({ commit }, nickName) {
      commit('SET_NICKNAME', { nickName: nickName })
    }
  }
}

export default user
