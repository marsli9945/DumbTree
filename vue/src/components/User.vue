<template>
    <div>
        User
    </div>
</template>

<script>
export default {
    name: 'User',
    mounted() {
        console.log(this.$route.params.id)
    }
}
</script>

