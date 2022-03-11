<template>
  <div class="container text-center pt-5">
    <h4>
      Welcome
      {{ ($auth.user) ? $auth.user.name : null }}
      to Authentication App
    </h4>
  </div>
</template>