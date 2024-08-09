<template lang="">
   <div class="flex mb-4 items-center bg-slate-300  p-4 rounded-2xl">
                <p class="w-full text-grey-darkest">Задача: {{task.name}} Дата: {{task.data}} Время создания: {{task.data_create}} </p>
                <button @click="done()" class="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded bg-green-500 rounded-lg hover:text-white text-green border-green hover:bg-green">Выполнено</button>
                <button @click="del()" class="flex-no-shrink p-2 ml-2 border-2 bg-red-500 rounded text-red border-red rounded-lg hover:text-white hover:bg-red">Удалить</button>
    </div>
</template>
<script>
import { Del,Done,Get } from '../../../wailsjs/go/main/App';

export default {
    name,
    props: ['task'],
    methods: {
        del() {
        let del_pass = confirm("Вы действительно хотите удалить?");
        if (del_pass == true) {
            Del(this.task.id).then(result => {
            this.$emit('test')
         })
        }
        },
        done() {
            Done(this.task.id).then(result => {
            this.$emit('test')
         })
        }
    },
}
</script>
<style lang="">
    
</style>