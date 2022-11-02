<template>
  <div class="target">
    <h1>Current Targets:</h1>

    <table class="table">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Id</th>
        <th scope="col">Type</th>
        <th scope="col">Source</th>
        <th scope="col">Frequency</th>
        <th scope="col">Unit</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(target, index) in targets">
        <th scope="row"> {{index}} </th>
        <td>{{target.Id}}</td>
        <td>{{target.type}}</td>
        <td>{{target.source}}</td>
        <td>{{target.frequency}}</td>
        <td>{{target.unit}}</td>
        {{target}}
      </tr>
      </tbody>
    </table>
  </div>


  <div class="container py-4 px-3 mx-auto">
    <h1>Target Vue!</h1>
    <button class="btn btn-primary">Primary button</button>
  </div>

</template>

<script>
export default {
  data() {
    return {
      targets: []
    };
  },
  async created() {
    const response = await fetch(`http://localhost:1323/api/target`);
    this.targets = await response.json();
  }
}
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
