<template>
  <div class="alert alert-primary" role="alert">
    <ol>
      <template v-for="(error, index) in errors">
        <li :key="index">
          {{ error }}
          <br />
        </li>
      </template>
    </ol>
  </div>
</template>

<script>
export default {
  data() {
    return {
      errors: [],
    }
  },
}
</script>
