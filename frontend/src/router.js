import { createWebHistory, createRouter } from 'vue-router'
import todo from './pages/todo.vue'
import Todo_done from './pages/todo_done.vue'
const routes = [
    { path: "/", component: todo, props: true },
    { path: "/tododone", component: Todo_done },
  ]
  
  export default new createRouter({
    mode: history,
    history: createWebHistory(),
    routes
  })