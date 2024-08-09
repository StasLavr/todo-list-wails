<script>
import navbar from './components/navbar.vue'
import task from './components/task.vue'
import task2 from './components/task2.vue'
import {Add,Get} from './../../wailsjs/go/main/App'

export default {
    components: {
        task,
        task2,
        navbar,
    },
    data() {
        return {
            name: "",
            data: "",
            info: "",
        }
    },
    methods: {
        Adds() {
         if (this.name.trim().length < 2 || this.data.trim().length < 2) {
         alert('Пожалуйста, заполните все поля не менее чем 2 символами');
        return;
      } else {
        Add(this.name,this.data).then(result => {
        this.update()
     })
    }},
    update() {
        Get().then(result => {
        this.info = result
     })
    },
    },
    beforeMount() {
        Get().then(result => {
        this.info = result
     })
    },

}

</script>

<template>
  <navbar/>
<div class="bg-neutral-800 w-screen min-h-screen flex justify-center items-center">
   <div class="min-h-80 w-full flex items-center justify-center bg-teal-lightest font-sans">
	<div class="bg-white rounded shadow p-6 m-4 w-full lg:max-w-6xl lg:max-w-lg">
        <div class="mb-4">
            <h1 class="text-grey-darkest">Todo List</h1>
            <div class="flex mt-4">
                <input v-model="this.name" class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker rounded-lg" placeholder="Задача">
                <input v-model="this.data" class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-grey-darker rounded-lg" placeholder="Дата">
                <button @click="Adds()"  class="flex-no-shrink p-2 ml-2 border-2 rounded bg-blue-500 text-red border-red rounded-lg hover:text-white hover:bg-red">Добавить</button>
            </div>
        </div>  
        <div>
            <task v-for="west in this.info" :task="west" @test="update()"></task>
        </div>
    </div>
    </div>
</div>
</template>

<style>
</style>