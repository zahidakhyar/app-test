<template>
  <div class="container">
    <div class="row align-items-center g-0" style="height: 80vh">
      <div class="col-12">
        <div class="card m-auto" style="max-width: 300px">
          <div class="card-header">
            <h4 class="card-title text-center">Profile</h4>
          </div>
          <div class="card-body">
            <b-form-group label="Name" label-for="name">
              <b-form-input
                id="name"
                type="text"
                :value="$auth.user.name"
                readonly
              />
            </b-form-group>
            <b-form-group label="Email" label-for="email">
              <b-form-input
                id="email"
                type="email"
                :value="$auth.user.email"
                readonly
              />
            </b-form-group>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  auth: 'auth',
}
</script>
