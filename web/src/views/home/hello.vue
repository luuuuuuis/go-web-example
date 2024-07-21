<template>
    <body>
        <div>
            <button @click="getData">点击获取后端数据</button>
            <p >{{ datas.message }}</p>
        </div>
    </body>

</template>


<script>
import url from "../common/url";
import httpClient from '../../utils/request';

    export default{
        data(){
            return{
                datas: [],
                helloUrl: url.helloUrl
            }
        },
        methods:{
            getData(){
                httpClient.get(this.helloUrl)
                .then(reponse => {
                    console.log(reponse)
                    this.datas = reponse
                })
                .catch()
            }
        }
    }
</script>