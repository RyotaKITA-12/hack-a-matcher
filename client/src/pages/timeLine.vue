<template>
    <main class="form-profile text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal"><b>Home</b></h1>
        <hr>
        <br>
        <button class="w=30 btn btn-success" @click="openModal">投稿</button>
        <div id="overlay" v-show="showContent">
            <form @submit.prevent="submit">
                <div id="content" class="rounded">
                    <button class="btn-circle btn-outline-dark " @click="closeModal">x</button>
                    <h1 class="h1 mb-3 fw-normal text-success"><b>募集</b></h1>
                    <hr>
                    <br>
                    <p class="h4">
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" v-model="post_type" name="post_type"
                                value="dev">
                            <label class="form-check-label text-secondary" for="dev">開発</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input" type="radio" v-model="post_type" name="post_type"
                                value="study">
                            <label class="form-check-label text-secondary" for="study">勉強会</label>
                        </div>
                    </p>
                    <br>
                    <div class="input-container">
                        <label for="title" class="text-secondary"><b>タイトル</b></label>
                        <input id="title" v-model="title" class="form-control" placeholder="Title" required>
                    </div>
                    <br>
                    <br>
                    <div class="input-container">
                        <label for="skills" class="text-secondary form-label"><b>スキル</b></label>
                        <Multiselect id="skills" v-model="value" mode="tags" placeholder="Select skills"
                            :options="options" :searchable="true" :createTag="true" style="width: 200px;" />
                    </div>
                    <br>
                    <br>
                    <div class="input-container">
                        <label for="recruit_name" class="text-secondary"><b>募集人数</b></label>
                        <input id="recruit_name" v-model="recruit_num" class="form-control" placeholder="100" required>
                    </div>
                    <br>
                    <br>
                    <div class="input-container">
                        <label for="view_period" class="text-secondary"><b>掲載期限</b></label>
                        <input id="view_period" type="date" v-model="view_period"
                            :min="new Date().toISOString().split('T')[0]">
                    </div>
                    <br>
                    <br>
                    <div class="input-container">
                        <label for="FormControlTextarea" class="text-secondary form-label"><b>本文</b></label>
                        <textarea id="FormControlTextarea" class="form-control" v-model="content" placeholder="本文"
                            rows="6"></textarea>
                    </div>
                    <br>
                    <br>
                    <br>
                    <button class="w-50 btn btn-lg btn-success" type="submit">投稿する</button>
                    <br>
                    <br>
                </div>
            </form>
        </div>
        <br>
        <br>
        <div class="profile-container">
            <div class="overflow-scroll bg-dark border border-dark border-5" style="height:600px; padding: 50px;">
                <div v-for="item in posts">
                    <div class="bg-light rounded border border-4" style="padding-right: 20px;padding-left: 20px;">
                        <br>
                        <p>投稿者 : {{ item.user_name }}　　募集人数 : {{ item.recruit_num }}</p>
                        <h1 class="h6 text-dark"><b>{{ item.title }}</b></h1>
                        <hr>
                        <br>
                        <p>{{ item.content }}</p>
                        <br>
                        <hr>
                        <p>{{ item.CreatedAt.substr(0, 10) }}　〜　{{ item.view_period.substr(0, 10) }}</p>
                    </div>
                    <br>
                </div>
            </div>
        </div>
    </main>
</template>
<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"
import Multiselect from '@vueform/multiselect'
/* import moment from "moment" */

export default {
    name: "timeline",
    components: { Multiselect },
    setup() {
        const store = useStore()
        const posts = ref([])

        var showContent = ref(false)
        const openModal = () => {
            showContent.value = true
        }
        const closeModal = () => {
            showContent.value = false
        }

        const value = []
        const options = ['js', 'node', 'vue', 'react', 'typescript']

        onMounted(async () => {
            try {
                const { data } = await axios.get('post')
                console.log(data)
                data.forEach(elem => {
                    posts.value.push(elem)
                })
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        const title = ref('');
        const recruit_num = ref('');
        const view_period = ref('');
        const content = ref('');


        const submit = async () => {
            const { data } = await axios.get('user')
            await axios.post('register/post', {
                user_id: `${data.ID}`,
                user_name: `${data.user_name}`,
                title: title.value,
                recruit_num: recruit_num.value,
                content: content.value,
                view_period: view_period.value,
            })
            showContent.value = false
        }

        return {
            posts,
            showContent,
            openModal,
            closeModal,
            title,
            recruit_num,
            content,
            view_period,
            submit,
            value,
            options
        }
    }
}
</script>
<style>
.form-profile {
    width: 100%;
    max-width: 500px;
    padding: 15px;
    margin: auto;
}

.form-profile .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-profile .form-control:focus {
    z-index: 2;
}

/* .profile-container div {       */
/*     background-color: #F0F0F5; */
/* }                              */

.container p {
    display: inline;
    text-align: left;
}

.profile-container h2 {
    text-align: left;
}

.profile-rect svg {
    margin-left: 5px;
    margin-right: 5px;
}

.rect-left {
    float: left;
}

.bg-gray {
    background-color: #F5F5FF;
}

#overlay {
    z-index: 1;

    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);

    display: flex;
    align-items: center;
    justify-content: center;

}

#content {
    z-index: 2;
    width: 450px;
    padding: 1em;
    background: #fff;
}

input[type="date"] {
    padding: 0 10px;
    width: 200px;
    height: 40px;
    border: solid 1px #CCC;
    box-sizing: border-box;
    border-radius: 5px;
    font-size: 15px;
    color: #999;
    box-shadow: 0px;
}

.btn-circle {
    font-size: 100%;
    font-weight: bold;
    border: 1px solid #999;
    color: #999;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 100%;
    width: 1.5em;
    line-height: 1.3em;
    cursor: pointer;
    transition: .2s;
}
</style>
<style src="@vueform/multiselect/themes/default.css">
</style>
