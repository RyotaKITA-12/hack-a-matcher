<template>
    <main class="form-profile text-center">
        <br>
        <h1 class="h1 mb-3 fw-normal"><b>プロフィール</b></h1>
        <hr>
        <br>
        <div class="overflow-scroll" style="height:750px;">
            <div class="profile-container">
                <h2 class="h6" style="float:right; textalign: left; color:#2277FF;">
                    <a hlef="#" @click.stop.prevent="openEditProfile"
                        style="float:right; textalign: left;"><b>編集</b></a>
                </h2>
                <div id="overlay" v-show="showEditProfile">
                    <form @submit.prevent="submit">
                        <div id="content" class="rounded">
                            <button class="btn-circle btn-outline-dark " @click="closeEditProfile">x</button>
                            <h1 class="h1 mb-3 fw-normal text-primary"><b>個人情報</b></h1>
                            <hr>
                            <br>
                            <div class="input-container">
                                <label for="birth" class="text-secondary"><b>生年月日</b></label>
                                <input id="birth" type="date" v-model="birth"
                                    :max="new Date().toISOString().split('T')[0]">
                            </div>
                            <br>
                            <br>
                            <div class="input-container">
                                <label for="gender" class="text-secondary"><b>性別</b></label>
                                <select id="gender" v-model="gender" style="width: 200px;" class="form-control"
                                    required>
                                    <option>男性</option>
                                    <option>女性</option>
                                    <option>その他</option>
                                </select>
                            </div>
                            <br>
                            <br>
                            <div class="input-container">
                                <label for="address" class="text-secondary"><b>住所</b></label>
                                <select id="address" v-model="address" style="width: 200px;" class="form-control"
                                    required>
                                    <option>北海道</option>
                                    <option>青森県</option>
                                    <option>岩手県</option>
                                    <option>宮城県</option>
                                    <option>秋田県</option>
                                    <option>山形県</option>
                                    <option>福島県</option>
                                    <option>茨城県</option>
                                    <option>栃木県</option>
                                    <option>群馬県</option>
                                    <option>埼玉県</option>
                                    <option>千葉県</option>
                                    <option>東京都</option>
                                    <option>神奈川県</option>
                                    <option>新潟県</option>
                                    <option>富山県</option>
                                    <option>石川県</option>
                                    <option>福井県</option>
                                    <option>山梨県</option>
                                    <option>長野県</option>
                                    <option>岐阜県</option>
                                    <option>静岡県</option>
                                    <option>愛知県</option>
                                    <option>三重県</option>
                                    <option>滋賀県</option>
                                    <option>京都府</option>
                                    <option>大阪府</option>
                                    <option>兵庫県</option>
                                    <option>奈良県</option>
                                    <option>和歌山県</option>
                                    <option>鳥取県</option>
                                    <option>島根県</option>
                                    <option>岡山県</option>
                                    <option>広島県</option>
                                    <option>山口県</option>
                                    <option>徳島県</option>
                                    <option>香川県</option>
                                    <option>愛媛県</option>
                                    <option>高知県</option>
                                    <option>福岡県</option>
                                    <option>佐賀県</option>
                                    <option>長崎県</option>
                                    <option>熊本県</option>
                                    <option>大分県</option>
                                    <option>宮崎県</option>
                                    <option>鹿児島県</option>
                                    <option>沖縄県</option>
                                </select>
                            </div>
                            <br>
                            <br>
                            <br>
                            <button class="w-50 btn btn-lg btn-primary" type="submit">更新</button>
                            <br>
                            <br>
                        </div>
                    </form>
                </div>
                <h2 class="h6"><b>個人情報</b></h2>
                <div class="container bg-gray rounded">
                    <div class="row">
                        <p></p>
                        <p class="text-secondary col-4">Eメール</p>
                        <p class="text-dark col-8">{{ email }}</p>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">生年月日</p>
                        <p class="text-dark col-8">{{ birth }}</p>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">性別</p>
                        <p class="text-dark col-8">{{ gender }}</p>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">住所</p>
                        <p class="text-dark col-8">{{ address }}</p>
                    </div>
                </div>
            </div>
            <br>
            <div class="profile-container">
                <h2 class="h6" style="float:right; textalign: left; color:#2277FF;">
                    <a hlef="#" @click.stop.prevent="openEditProfile"
                        style="float:right; textalign: left;"><b>編集</b></a>
                </h2>
                <h2 class="h6"><b>リンク</b></h2>
                <div class="container bg-gray rounded">
                    <div class="row">
                        <p></p>
                        <p class="text-secondary col-4">Portfolio</p>
                        <p class="text-dark col-8"><a v-bind:href="`https://${portfolio}`">{{ portfolio }}</a></p>
                    </div>
                    <div class="row">
                        <p></p>
                        <p class="text-secondary col-4">Twitter</p>
                        <p class="text-dark col-8"><a v-bind:href="`https://twitter.com/${twitter}`">{{ twitter }}</a>
                        </p>
                    </div>
                    <div class="row">
                        <p></p>
                        <p class="text-secondary col-4">Github</p>
                        <p class="text-dark col-8"><a v-bind:href="`https://github.com/${github}`">{{ github }}</a></p>
                    </div>
                </div>
            </div>
            <br>
            <div class="profile-container">
                <h2 class="h6" style="float:right; textalign: left; color:#2277FF;">
                    <a hlef="#" @click.stop.prevent="openEditProfile"
                        style="float:right; textalign: left;"><b>編集</b></a>
                </h2>
                <h2 class="h6"><b>スキル</b></h2>
                <div class="container  bg-gray rounded">
                    <div class="row">
                        <p></p>
                        <p class="text-secondary col-4">役割</p>
                    </div>
                    <div class="profile-rect">
                        <svg class="rounded-pill rect-left" width="100" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                dominant-baseline="central">エンジニア</text>
                        </svg>
                        <svg class="rounded-pill rect-left" width="100" height="35">
                            <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                            <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                dominant-baseline="central">デザイナー</text>
                        </svg>
                    </div>
                    <br>
                    <br>
                    <hr>
                    <br>
                    <div class="row">
                        <p class="text-secondary col-4">使用経験</p>
                        <p class="text-dark col-8">趣味</p>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">言語</p>
                        <div class="profile-rect">
                            <svg class="rounded-pill rect-left" width="130" height="35">
                                <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                                <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                    dominant-baseline="central">Python</text>
                            </svg>
                            <svg class="rounded-pill rect-left" width="130" height="35">
                                <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                                <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                    dominant-baseline="central">Go</text>
                            </svg>
                            <svg class="rounded-pill rect-left" width="130" height="35">
                                <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                                <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                    dominant-baseline="central">JavaScript</text>
                            </svg>
                        </div>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">使用経験</p>
                        <p class="text-dark col-8">授業</p>
                    </div>
                    <div class="row">
                        <p class="text-secondary col-4">言語</p>
                        <div class="profile-rect">
                            <svg class="rounded-pill rect-left" width="100" height="35">
                                <rect width="100%" height="100%" fill="#E7E7E7"></rect>
                                <text x="50%" y="50%" fill="#334" text-anchor="middle"
                                    dominant-baseline="central">Pyhon</text>
                            </svg>
                        </div>
                    </div>
                    <br>
                    <hr>
                    <br>
                    <br>
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
import moment from "moment"


export default {
    name: "profile",
    setup() {
        const email = ref('none')
        const birth = ref('none')
        const gender = ref('none')
        const address = ref('none')
        const portfolio = ref('none')
        const twitter = ref('none')
        const github = ref('none')
        const store = useStore()

        onMounted(async () => {
            try {
                const p = await axios.get('profile')
                const l = await axios.get('link')
                email.value = `${p.data.email}`
                birth.value = `${p.data.birthday}`
                birth.value = moment(birth.value).format("YYYY年MM月DD日")
                gender.value = `${p.data.gender}`
                address.value = `${p.data.address}`
                portfolio.value = `${l.data.portfolio}`
                twitter.value = `${l.data.twitter}`
                github.value = `${l.data.github}`
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        var showEditProfile = ref(false)
        const openEditProfile = () => {
            showEditProfile.value = true
        }
        const closeEditProfile = () => {
            showEditProfile.value = false
        }

        const submit = async () => {
            const { data } = await axios.get('user')
            console.log(data)
            await axios.post('register/profile', {
                user_id: `${data.ID}`,
                email: `${data.email}`,
                birthday: birth.value,
                gender: gender.value,
                address: address.value,
            })
            showEditProfile.value = false
        }

        return {
            email,
            birth,
            gender,
            address,
            portfolio,
            twitter,
            github,
            submit,
            showEditProfile,
            openEditProfile,
            closeEditProfile
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
    margin: 5px;
}

.rect-left {
    float: left;
}

.bg-gray {
    background-color: #F5F5FF;
}

a:hover {
    cursor: pointer;
}
</style>
