<template>
    <div>
        <h1>Hello Vue</h1>
        <p>
            {{ message }}
        </p>
        <button @click="changeMessage">Change Message</button>
    </div>

</template>

<script>
export default {
    data() {
        return {
            message: 'Hello Vue.js'
        }
    },
    methods: {
        changeMessage() {
            this.message = 'Hello Vue.js2'
        }
    }

}

</script>

<style>

</style>