<template>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark">
        <div class="container-fluid">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <router-link to="/" class="nav-link">
                        <a href="#" class="navbar-brand"><img src="./img/hack_a_matcher.png" height="30"
                                alt="hack-a-matcher"></a>
                    </router-link>
                </li>
            </ul>
            <ul class="navbar-nav my-2 my-lg-0 ms-auto">
                <template v-if="auth">
                    <li class="nav-item">
                        <router-link to="/login" class="nav-link" @click="logout">
                            <button class="btn btn-outline-danger">ログアウト</button>
                        </router-link>
                    </li>
                </template>
                <template v-if="!auth">
                    <li class="nav-item">
                        <router-link to="/login" class="nav-link">
                            <button class="btn btn-outline-primary">ログイン</button>
                        </router-link>
                    </li>
                    <li class="nav-item">
                        <router-link to="/register" class="nav-link">
                            <button class="btn btn-outline-success">新規登録</button>
                        </router-link>
                    </li>
                </template>
            </ul>
        </div>
    </nav>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import axios from 'axios'

export default {
    name: "Nav",
    setup() {
        const store = useStore()
        const router = useRouter()
        const auth = computed(() => store.state.auth)
        const logout = async () => {
            await axios.post('logout', {})
            store.dispatch('setAuth', false)
            await router.push('/login')
        }
        return { auth, logout }
    }
}
</script>
