<template>
  <div>
    <b-navbar type="is-warning">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img :src="logo" />
        </b-navbar-item>
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ name: 'Vehiculos' }"
          ><b-icon icon="car" class="mr-1"></b-icon>Vehículos</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'ReporteVehiculos' }"
          ><b-icon icon="table-check" class="mr-1"></b-icon>
          Reporte</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'Escritorio' }"
          ><b-icon icon="monitor-dashboard" class="mr-1"></b-icon>
          Escritorio</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'Ajustes' }"
          ><b-icon icon="cog-outline" class="mr-1"></b-icon>
          Ajustes</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'AcercaDe' }"
          ><b-icon icon="information" class="mr-1"></b-icon> Acerca
          de</b-navbar-item
        >
      </template>

      <template #end>
        <b-navbar-item tag="div">
          <div class="buttons">
            <a href="https://parzibyte.me/#contacto" class="button is-danger">
              <b-icon icon="help"></b-icon>
              <strong>Ayuda y soporte</strong>
            </a>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
    <section class="section">
      <router-view></router-view>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong
            >Programa para el control de estacionamientos / parqueaderos </strong
          >creado y mantenido con
          <b-icon icon="heart" type="is-danger"></b-icon> por
          <a href="https://parzibyte.me/blog">Parzibyte</a>
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Logo from "./assets/parzibyte.png";
export default {
  name: "app",
  data: () => ({
    logo: Logo,
  }),
  methods: {},
};
</script>