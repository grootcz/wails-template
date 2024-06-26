<template>
  <div class="home-container">
    <div class="home-menu">
      <div class="menu-empty"/>

      <div v-for="(button, index) in buttons" :key="index" class="tooltip">
        <a href="#" @click.prevent="menuButtonClicked(index)" :style="{ color: button.color }">
          <i aria-hidden="true">
            <font-awesome-icon :icon="button.icon" />
          </i>
          <span class="tooltiptext">{{ button.description }}</span>
        </a>
        <div class="menu-empty"/>
      </div>

    </div>

    <div class="home-content">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useStore } from 'vuex'
import {useRouter} from "vue-router"

const store = useStore()
const result = ref("Please enter your name below 👇")
const name = ref("")
const router  = useRouter();

const greet = () => {
  window.go.main.App.Greet(name.value).then(response => {
    // store.commit('increment')
    // result.value = `${store.state.count}. ${response}`
  })
}

interface Button {
  color: string;
  description: string;
  icon: [string, string];
  path: string;
}

const buttons = ref<Button[]>([
  { color: '#e3e3e3', description: 'Button1 Description', icon:['fas', 'rectangle-list'], path:"/home/menu1",},
  { color: '#e3e3e3', description: 'Button2 Description', icon:['fas', 'coffee'], path:"/home/menu2",},
  { color: '#e3e3e3', description: 'Button3 Description', icon:['fas', 'apple-alt'], path:"/home/menu3",},
  { color: '#e3e3e3', description: 'Button4 Description', icon:['fas', 'bell'], path:"/home/menu4",},
]);

const menuButtonClicked = (index: number) => {
  router.push(buttons.value[index].path)

  buttons.value.forEach(button => {
    button.color = "#e3e3e3"
  })
  buttons.value[index].color = "#3eb0e0"

  store.commit("home/changeName", index)
};

// 在组件挂载时执行的初始化函数
const initialize = () => {
  buttons.value.forEach(button => {
    button.color = "#e3e3e3"
  })

  const index = store.state.home.menuButtonIndex
  buttons.value[index].color = "#3eb0e0";
  router.push(buttons.value[index].path)
};

// 在组件卸载前执行的清理函数
const cleanup = () => {
  // 在此处执行你的清理逻辑
};

// 使用 onMounted 生命周期钩子
onMounted(() => {
  initialize();
  return
});

// 使用 onBeforeUnmount 生命周期钩子
onBeforeUnmount(() => {
  cleanup();
  return
});
</script>

<style lang="scss" scoped>
.home-container {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: flex-start;
  align-items: center;
  height: 100%;
  width: 100%;
  //background: transparent;
  background-color: rgba(0, 0, 0, 0);
}

.home-menu{
  height: 100%;
  width: 47px;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  justify-content: flex-start;
  align-items: center;
}

.home-content{
  flex: 1;
  height: 100%;
  background-color: rgba(0, 0, 0, 0);
}

a {
  position: relative;
  display: block;
  width: 28px;
  height: 28px;
  text-align: center;
  line-height: 34px;
  background: #000000;
  border-radius: 50%;
  font-size: 14px;
  color: #e3e3e3;
  transition: .5s;
}

a:before {
  content: '';
  position: absolute;
  top:0;
  left:0;
  width:100%;
  height:100%;
  border-radius:50%;
  background: #d35400;
  transition: .5s;
  transform: scale(.9);
  z-index: -1;
}

a:hover:before {
  transform: scale(1.2);
  box-shadow: 0 0 10px #d35400;
  filter: blur(3px);
}

a:hover {
  color: #ffaf02;
  box-shadow: 0 0 10px #d35400;
  text-shadow: 0 0 10px #d35400;
}

.tooltip .tooltiptext {
  visibility: hidden;
  width: 170px;
  height: 27px;
  background-color: rgba(0, 0, 0, 0);
  color: #fff;
  text-align: left;
  border-radius: 6px;
  padding: 5px;
  position: absolute;
  z-index: 1;
  top: 50%;
  left: 150%; /* Position the tooltip to the right */
  margin-top: -20px; /* Use negative margin to center the tooltip vertically */
  opacity: 0;
  transition: opacity 0.3s;
}

.tooltip:hover .tooltiptext {
  visibility: visible;
  opacity: 1;
}

.menu-empty{
  width: 28px;
  height: 17px;
  background-color: rgba(0, 0, 0, 0);
}
</style>
