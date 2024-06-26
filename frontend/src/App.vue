<template>
  <div class="world-all">
    <div class="header-all" height="49px">
      <Header/>
    </div>
    <div class="body-all">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Header from "./components/Header.vue";

</script>

<style lang="scss">
@font-face {
  font-family: "Nunito";
  font-style: normal;
  font-weight: 400;
  src: local(""),
  url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
}

html {
  width: 100%;
  height: 100%;
  background-color: rgba(255, 0, 0, 0);
  color: #3e7fd5;
  font-family: "Nunito", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
  "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
  sans-serif;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
}

#app {
  width: 100%;
  height: 100%;
}

.world-all {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  justify-content: flex-start;
  align-items: center;
}

.header-all{
  height: 49px;
  width: 100%;
  padding: 0;
  margin: 0;
}

.body-all{
  flex: 1; /* 使 el-main 填充剩余高度 */
  width: 100%;
  padding: 0;
  margin: 0;
}

:root {
  --color-bg-primary: #d0d6df;
  --color-bg-primary-offset: #f1f3f7;
  --color-bg-secondary: #fff;
  --color-text-primary: #3a3c42;
  --color-text-primary-offset: #898c94;
  --color-orange: #dc9960;
  --color-green: #1eb8b1;
  --color-purple: #657cc4;
  --color-black: var(--color-text-primary);
  --color-red: #d92027;
}

</style>
