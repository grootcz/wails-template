<template>
  <div class="header-all" style="--wails-draggable:drag" >
    <div class="icon-box">
      <el-button style="width: 48px; height: 30px" @click="iconClicked">
        <div class="app-icon-show" style="width: 48px; height: 30px">
          <el-image class="logo" :src="logoImg" style="width: 48px; height: 30px" />
        </div>
      </el-button>
    </div>

    <div class="table-box">
      <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          style="height: 49px; width: 100%"
          background-color="rgba(0, 0, 0, 0.2)"
          @select="exampleClicked"
          text-color="rgba(255, 255, 255, 0.8)"
          :router="true"
      >
        <el-menu-item index="1" route="/home">主页</el-menu-item>
        <el-menu-item index="2" route="/features">Info</el-menu-item>
        <el-menu-item index="3">Info</el-menu-item>
        <el-menu-item index="4">Orders</el-menu-item>
      </el-menu>
    </div>

    <div class="user-box">
      <el-badge is-dot class="item" type="warning" :offset="[-2, 5]" badge-style="width: 6px; height: 6px;">
        <el-button style="width: 27px; height: 27px;" color="rgba(0, 0, 0, 0.2)" class="custom-icon-button" circle>
          <font-awesome-icon :icon="['fas', 'bell']" />
        </el-button>
      </el-badge>

      <el-button style="width: 27px; height: 27px;" color="rgba(0, 0, 0, 0.2)" class="custom-icon-button" circle>
        <font-awesome-icon :icon="['fas', 'rectangle-list']" />
      </el-button>

      <div class="user-head" style="width: 37px; height: 37px">
        <el-dropdown trigger="click" style="width: 37px; height: 37px">
          <el-button class="user-dropdown-btn" style="width: 37px; height: 37px" circle>
            <div class="user-head-show" style="width: 37px; height: 42px">
              <el-image class="logo" :src="logoImg" style="width: 37px; height: 37px" />
            </div>
          </el-button>

          <template #dropdown>
            <el-dropdown-menu class="menu-all">
              <li class="menu-item">
                <button class="menu-button">
                  <font-awesome-icon :icon="['fas', 'rectangle-list']" />
                  个人设置
                </button>
              </li>

              <li class="menu-item">
                <button class="menu-button">
                  <font-awesome-icon :icon="['fas', 'rectangle-list']" />
                  个人设置
                </button>
              </li>

              <li class="menu-item">
                <button class="menu-button">
                  <font-awesome-icon :icon="['fas', 'rectangle-list']" />
                  软件设置
                </button>
              </li>

              <li class="menu-item">
                <button class="menu-button">
                  <font-awesome-icon :icon="['fas', 'rectangle-list']" />
                  关于软件
                </button>
              </li>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>

    <div class="control-box">
      <el-button @click="smallClicked" style="width: 17px; height: 17px;" color="rgba(255, 223, 92, 0.5)" :icon="SemiSelect" size="small" circle />
      <el-button @click="maxClicked" style="width: 17px; height: 17px;" color="rgba(122, 255, 91, 0.5)" :icon="DocumentCopy" size="small" circle />
      <el-button @click="closeClicked" style="width: 17px; height: 17px;" color="rgba(255, 78, 78, 0.5)" :icon="CloseBold" circle />
      <div class="control-empty-right"/>
    </div>
  </div>
</template>

<script setup lang="ts">
// This template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref, onMounted } from 'vue'
import { useStore } from 'vuex'
import { DocumentCopy, CloseBold, SemiSelect, Message, Finished, Setting, MagicStick, Operation  } from '@element-plus/icons-vue'
import logoImg from '../assets/images/github.png'
import {useRouter} from "vue-router"

const store = useStore()
const name = ref("")
const activeIndex = "1"
const router  = useRouter();

// 在组件挂载时执行的初始化函数
const initialize = () => {
  router.push({ path:'/home' })
};

// 使用 onMounted 生命周期钩子
onMounted(() => {
  initialize();
  return
});

const greet = () => {
  window.go.main.App.Greet(name.value).then(response => {
    // store.commit('increment')
    // result.value = `${store.state.count}. ${response}`
  })
}

const closeClicked = () => {
  window.runtime.Quit()
}

const maxClicked = () => {
  const isMaxObj = window.runtime.WindowIsMaximised()
  isMaxObj.then(isMaxInfo =>{
    if (isMaxInfo.valueOf() == true) {
      window.runtime.WindowUnmaximise()
    } else {
      window.runtime.WindowMaximise()
    }
  })
}

const smallClicked = () => {
  window.runtime.WindowMinimise()
}

const exampleClicked = () => {

}

const iconClicked = () => {
  window.runtime.BrowserOpenURL("https://www.baidu.com")
}
</script>

<style lang="scss" scoped>
.header-all {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-between;
  align-items: center;
  height: 49px;
  width: 100%;
  background-color: rgba(0, 0, 0, 0.3);
}

.icon-box {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
  align-items: center;
  height: 49px;
  width: 170px;
  //background: transparent;
  background-color: rgba(0, 0, 0, 0);
}

.table-box {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
  align-items: center;
  height: 49px;
  width: 670px;
  background: transparent;
}

.user-box {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
  align-items: center;
  height: 49px;
  width: 180px;
  background: transparent;
}

.control-box {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
  align-items: center;
  height: 49px;
  width: 120px;
  background: transparent;
}

.control-empty-right {
  height: 49px;
  width: 10px;
  background: transparent;
}

.custom-icon-button >>> .svg-inline--fa {
  font-size: 30px; /* 自定义图标大小 */
}

.app-icon-show{
  display: -webkit-flex;
  /* Safari */
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

.user-head-show {
  display: -webkit-flex;
  /* Safari */
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

.user-dropdown-btn {
  color: #fff;
  background-color: #26272c;
  border-color: #26272c;
  display: -webkit-flex;
  /* Safari */
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

.user-dropdown-btn:hover {
  background-color: #26272c;
  border-color: #26272c;
}

.user-dropdown-btn:focus {
  background-color: #7997f3;
  border-color: #7997f3;
}

.menu-all {
  width: 170px;
  background-color: rgba(0, 0, 0, 0.2);
  display: -webkit-flex;
  /* Safari */
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;
  top: 0;
  left: 0;
  padding: 0;
  margin: 0;
}

.el-menu--horizontal.el-menu-demo {
  border-bottom: none;
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

.menu-button {
  font: inherit;
  border: 0;
  padding: 8px 8px;
  padding-right: 36px;
  //width: 100%;
  width: 170px;
  border-radius: 0px;
  text-align: left;
  display: flex;
  align-items: center;

  background-color: var(--color-bg-secondary);
  &:hover {
    background-color: var(--color-bg-primary-offset);
    & + .menu-sub-list {
      display: flex;
    }
    svg {
      stroke: var(--color-text-primary);
    }
  }
  svg {
    flex-shrink: 0;
    width: 20px;
    height: 20px;
    margin-right: 10px;
    stroke: var(--color-text-primary-offset);
    &:nth-of-type(2) {
      margin-right: 0;
      position: absolute;
      right: 8px;
    }
  }

  &--delete {
    &:hover {
      color: var(--color-red);
      svg:first-of-type {
        stroke: var(--color-red);
      }
    }
  }

  &--orange {
    svg:first-of-type {
      stroke: var(--color-orange);
    }
  }

  &--green {
    svg:first-of-type {
      stroke: var(--color-green);
    }
  }
  &--purple {
    svg:first-of-type {
      stroke: var(--color-purple);
    }
  }
  &--black {
    svg:first-of-type {
      stroke: var(--color-black);
    }
  }

  &--checked {
    svg:nth-of-type(2) {
      stroke: var(--color-purple);
    }
  }
}

</style>