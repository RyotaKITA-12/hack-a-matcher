<template>
    <div class="container text-center">
        <br>
        <br>
        <br>
        <h1 style="text-shadow: 1px 1px 5px #808080; color:#f8f8ff"><b>クリエイターをつなぐ<br>マッチングアプリ</b></h1>
        <br>
        <br>
        <img src="../components/img/hack_a_matcher_2.png" width="450">
        <br>
        <br>
        <br>
        <span v-html=message></span>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"

export default {
    name: "Index",
    setup() {
        const message = ref('<p class="h5"><b>ログインは<a href="/login" style="color:#58E;">こちら</b></a></p>')
        const store = useStore()

        onMounted(async () => {
            try {
                const _ = await axios.get('user')
                message.value = '<p class="h5"><b>ホームは<a href="/home" style="color:#58E;">こちら</b></a></p>'
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        return { message }
    }
}
</script>
