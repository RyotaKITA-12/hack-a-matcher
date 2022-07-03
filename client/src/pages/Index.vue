<template>
    <div class="container">
        <br>
        <h1>{{ message }}</h1>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"

export default {
    name: "Index",
    setup() {
        const message = ref('You are not logged in!')
        const store = useStore()

        onMounted(async () => {
            try {
                const { data } = await axios.get('user')
                message.value = `Hi, ${data.user_name}`
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        return { message }
    }
}
</script>
