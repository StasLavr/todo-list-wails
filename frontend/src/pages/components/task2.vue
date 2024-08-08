<template lang="">
   <div class="flex mb-4 items-center">
                <p class="w-full line-through text-green">Задача: {{task.id}} {{task.name}} Дата: {{task.data}} Время создания: {{task.data_create}}</p>
                <button @click="done_off()" class="flex-no-shrink p-2 ml-4 mr-2 border-2 rounded bg-slate-400 hover:text-white  rounded-lg text-grey border-grey hover:bg-grey">Не выполнено</button>
                <button @click="del()" class="flex-no-shrink p-2 ml-2 border-2 rounded bg-red-500 text-red border-red rounded-lg hover:text-white hover:bg-red">Удалить</button>
   </div>
</template>
<script>
import {Del, Done_off} from '../../../wailsjs/go/main/App';

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
        done_off() {
            Done_off(this.task.id).then(result => {
            this.$emit('test')
         })
        }
    },
}

</script>
<style lang="">
    
</style>