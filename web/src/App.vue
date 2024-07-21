<!-- 框架 -->
<template>
  <!-- 路由占位符，会导入匹配到的$route.path的视图组件 -->
  <RouterView></RouterView>
</template>

<!-- 交互 -->
<script>

</script>

<!-- 样式 -->
<style>
  html,body{
    width: 100%;
    height: 100%;
    padding: 0;
    margin: 0;
  }
  #nprogress .bar {
    /*自定义进度条颜色*/
    background: #2186c0 !important
  }
</style>