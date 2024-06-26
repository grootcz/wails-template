import { createStore } from 'vuex'

import home from './modules/home'
import global from './modules/global'

const store = createStore({
    // 公共模板直接在这里展开就可以
    ...global,
    modules: {
        home,
    }
})

export default store;