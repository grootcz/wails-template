
const home = {
    namespaced: true,
    state() {
        return {
            menuButtonIndex: 0,
        }
    },
    mutations: {
        changeName(state: any, menuButton: number): void {
            state.menuButtonIndex = menuButton;
        },
    },
    actions: {
        changeName({ commit }: any, menuButton: number): void {
            commit("changeName", menuButton);
        },
    },
};

export default home;