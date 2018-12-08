<template>
  <div id="app">
    <Navigation />
    <div class="container main">
      <GameBoard />
    </div>
    <div class="fixed-bottom">
      <div class="bg-dark p-4">
        <a href="https://github.com/decebal"><span class="fa fa-github"></span> Github</a>
      </div>
    </div>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.min.css';
import 'font-awesome/css/font-awesome.min.css';
import GameBoard from './components/Game.vue';
import Navigation from './components/Navigation.vue';

export default {
  name: 'app',
  components: {
    Navigation,
    GameBoard,
  },
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

body {
  padding-top: 50px;
}


* {
  margin: 0;
  padding: 0;
}
table {
  border-collapse: collapse;
  border-spacing: 0;
}

.main {
  margin-top: 25px;
}
</style>
