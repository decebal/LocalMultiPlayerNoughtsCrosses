<template>
  <div id="game">
    <div class="well well-lg" v-if="gameState.started" id="statusMessage">
      <h4>{{gameState.statusMessage}}</h4>
      <h5 v-if="gameState.started && !gameState.over">Your current symbol:
        <span v-if="playerNum === 0" style="color: red;">{{ playerSymbol }}</span>
        <span v-if="playerNum === 1" style="color: deepskyblue;">{{ playerSymbol }}</span>
      </h5>
      <button
        v-if="gameState.over"
        @click="startNew()"
        type="button"
        class="btn btn-success">
        Restart
      </button>
      <div style="padding-bottom: 15px;" ></div>
    </div>

    <div class="jumbotron animated" v-if="!gameState.started">
      <h1>{{gameState.statusMessage}}</h1>
      <div class="tt-cube-grid">
        <div class="tt-cube tt-cube1"></div>
        <div class="tt-cube tt-cube2"></div>
        <div class="tt-cube tt-cube3"></div>
        <div class="tt-cube tt-cube4"></div>
        <div class="tt-cube tt-cube5"></div>
        <div class="tt-cube tt-cube6"></div>
        <div class="tt-cube tt-cube7"></div>
        <div class="tt-cube tt-cube8"></div>
        <div class="tt-cube tt-cube9"></div>
      </div>
    </div>

    <div class="jumbotron animated" v-if="gameState.started">
      <Board />
    </div>
    <button
      v-if="!gameState.over"
      @click="startNew()"
      type="button"
      class="btn btn-warning">
      Give Up
    </button>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';
import Board from './Board.vue';

export default {
  name: 'game',
  components: { Board },
  computed: mapState([
    'gameState',
    'playerSymbol',
    'playerNum',
  ]),
  methods: mapActions([
    'startNew',
  ]),
};
</script>

<style scoped>
  #statusMessage{
    text-align: center;
  }

  .jumbotron{
    text-align: center;
    color: #334046;
  }

  .tt-cube-grid {
    width: 100px;
    height: 100px;
    margin: 100px auto;
  }

  .tt-cube-grid .tt-cube {
    width: 33%;
    height: 33%;
    background-color: white;
    float: left;
    -webkit-animation: tt-cubeGridScaleDelay 1.3s infinite ease-in-out;
    animation: tt-cubeGridScaleDelay 1.3s infinite ease-in-out;
  }

  .tt-cube-grid .tt-cube1 {
    -webkit-animation-delay: 0.2s;
    animation-delay: 0.2s; }
  .tt-cube-grid .tt-cube2 {
    -webkit-animation-delay: 0.3s;
    animation-delay: 0.3s; }
  .tt-cube-grid .tt-cube3 {
    -webkit-animation-delay: 0.4s;
    animation-delay: 0.4s; }
  .tt-cube-grid .tt-cube4 {
    -webkit-animation-delay: 0.1s;
    animation-delay: 0.1s; }
  .tt-cube-grid .tt-cube5 {
    -webkit-animation-delay: 0.2s;
    animation-delay: 0.2s; }
  .tt-cube-grid .tt-cube6 {
    -webkit-animation-delay: 0.3s;
    animation-delay: 0.3s; }
  .tt-cube-grid .tt-cube7 {
    -webkit-animation-delay: 0s;
    animation-delay: 0s; }
  .tt-cube-grid .tt-cube8 {
    -webkit-animation-delay: 0.1s;
    animation-delay: 0.1s; }
  .tt-cube-grid .tt-cube9 {
    -webkit-animation-delay: 0.2s;
    animation-delay: 0.2s; }

  @-webkit-keyframes tt-cubeGridScaleDelay {
    0%, 70%, 100% {
      -webkit-transform: scale3D(1, 1, 1);
      transform: scale3D(1, 1, 1);
    } 35% {
      -webkit-transform: scale3D(0, 0, 1);
      transform: scale3D(0, 0, 1);
    }
  }

  @keyframes tt-cubeGridScaleDelay {
    0%, 70%, 100% {
      -webkit-transform: scale3D(1, 1, 1);
      transform: scale3D(1, 1, 1);
    } 35% {
      -webkit-transform: scale3D(0, 0, 1);
      transform: scale3D(0, 0, 1);
    }
  }

  .animated {
    animation-duration: 0.8s;
    -webkit-animation-duration: 0.8s;
    -moz-animation-duration: 0.8s;
  }
</style>
