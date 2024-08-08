<script>
import navbar from './components/navbar.vue'
import task from './components/task.vue'
import task2 from './components/task2.vue'
import {Add,Get_done} from './../../wailsjs/go/main/App'


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
            user_id: "",
            info: "",
        }
    },
    methods: {
        Adds() {
        Add(1,this.name,this.data).then(result => {
        Get_done(1).then(result => {
        this.info = result
     })  
     })
    },
    update() {
        Get_done(1).then(result => {
        this.info = result
     })  
    }
    },
    beforeMount() {
        Get_done(1).then(result => {
        this.info = result
     })
    }
}

</script>

<template>
  <navbar/>
  <div class="bg-neutral-800 w-screen h-screen flex justify-center items-center">
    <!-- component -->
<div class="h-80 w-full flex items-center justify-center bg-teal-lightest font-sans">
	<div class="bg-white rounded shadow p-6 m-4 w-full lg:max-w-6xl lg:max-w-lg">
        <div class="mb-4">
            <h1 class="text-grey-darkest">Todo List Done</h1>
        </div>
        <div>
            <task2 v-for="west in this.info" :task="west" @test="update()"></task2>
        </div>
    </div>
</div>
  </div>
</template>

<style>
</style>