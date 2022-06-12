<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/omoshiro_go_moji/new">Create</router-link>
  </nav>
  <router-view />
</template>

<style lang="scss">
#app {
  text-align: center;
}

html {
  height: 100%;
}

body {
  height: 100%;
  margin-right: 300px;
  margin-left: 300px;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
