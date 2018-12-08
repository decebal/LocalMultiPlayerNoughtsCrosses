/* eslint-disable no-param-reassign, no-console */
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    gameState: {
      statusMessage: 'Connecting',
      started: false,
      over: false,
    },
    board: {
      row1: [{}, {}, {}],
      row2: [{}, {}, {}],
      row3: [{}, {}, {}],
    },
    playerNum: -1,
    playerSymbol: '',
  },

  mutations: {
    SOCKET_ONOPEN(state) {
      state.isConnected = true;
    },
    SOCKET_ONMESSAGE(state, gameState) {
      if (gameState.playerNum !== undefined) {
        state.playerNum = gameState.playerNum;
      } else {
        state.gameState.statusMessage = gameState.statusMessage;
        state.gameState.started = gameState.started;
        state.gameState.over = gameState.over;
        state.board.row1 = gameState.fields.slice(0, 3);
        state.board.row2 = gameState.fields.slice(3, 6);
        state.board.row3 = gameState.fields.slice(6, 9);
        state.playerSymbol = gameState.playerSymbols[state.playerNum];
      }
    },
    SOCKET_ONERROR(state, event) {
      console.error(state, event);
    },
    SOCKET_ONCLOSE(state, event) {
      console.error(state, event);
    },
  },

  actions: {
    makeMove(context, fieldNum) {
      Vue.prototype.$socket.send(fieldNum);
    },
    startNew() {
      Vue.prototype.$socket.send(10);
    },
  },
});
