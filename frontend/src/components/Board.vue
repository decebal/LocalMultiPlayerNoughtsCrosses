<template>
  <table id='board'>
    <tr>
      <td
        v-for="(field, index) in board.row1"
        @click="makeMove(index+0)"
        v-bind:key="index"
        class="top">
        <span v-if="field.set">{{field.symbol}}</span>
        <span v-if="!field.set" class="free-field"></span>
      </td>
    </tr>
    <tr>
      <td
        v-for="(field, index) in board.row2"
        @click="makeMove(index+3)"
        v-bind:key="index"
        class="middle">
        <span v-if="field.set">{{field.symbol}}</span>
        <span v-if="!field.set" class="free-field"></span>
      </td>
    </tr>
    <tr>
      <td
        v-for="(field, index) in board.row3"
        @click="makeMove(index+6)"
        v-bind:key="index"
        class="bottom">
        <span v-if="field.set">{{field.symbol}}</span>
        <span v-if="!field.set" class="free-field"></span>
      </td>
    </tr>
  </table>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: 'Board',
  computed: mapState([
    'board',
  ]),
  methods: mapActions([
    'makeMove',
  ]),
};
</script>

<style scoped>
  #board {
    padding: 0;
    margin: 0 auto;
  }

  .free-field {
    color: white;
  }

  #board tr td {
    width: 80px;
    height: 80px;
    border: 2px solid #b3b3b3;
    font-family: Helvetica, sans-serif;
    font-size: 30px;
    text-align: center;
    background: white;
  }

  #board tr td.top {
    border-top: none;
  }
  #board tr td.top:nth-child(1) {
    border-left: none;
  }
  #board tr td.top:nth-child(3) {
    border-right: none;
  }

  #board tr td.middle:nth-child(1) {
    border-left: none;
  }
  #board tr td.middle:nth-child(3) {
    border-right: none;
  }

  #board tr td.bottom {
    border-bottom: none;
  }
  #board tr td.bottom:nth-child(1) {
    border-left: none;
  }
  #board tr td.bottom:nth-child(3) {
    border-right: none;
  }

  #board tr td:hover {
    background: lightgreen;
    cursor: pointer;
  }
</style>
