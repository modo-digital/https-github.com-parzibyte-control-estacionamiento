<template>
  <div>
    <div class="columns">
      <b-modal
        :active.sync="mostrarModalCobrar"
        aria-modal
        aria-role="dialog"
        has-modal-card
        :destroy-on-hide="true"
        trap-focus
      >
        <Cobrar
          @cobrado="onCobroTerminado()"
          :vehiculo="vehiculoParaCobrar"
        ></Cobrar>
      </b-modal>
      <div class="column">
        <router-link
          :to="{ name: 'RegistrarVehiculo' }"
          class="button is-success mb-2"
        >
          Registrar nuevo
        </router-link>
        <b-field
          grouped
          message="Selecciona el rango de fechas para ver registros"
        >
          <b-field label="Inicio">
            <b-datepicker
              ref="seleccionadorFechaInicio"
              :append-to-body="true"
              v-model="fechaInicio"
              @input="onFechaInicioCambiada()"
              locale="es-MX"
              placeholder="Selecciona una fecha"
              :date-formatter="formateadorFecha"
              icon="calendar-today"
              close-on-click
            >
            </b-datepicker>
          </b-field>
          <b-field label="Fin">
            <b-datepicker
              :date-formatter="formateadorFecha"
              ref="seleccionadorFechaFin"
              :append-to-body="true"
              v-model="fechaFin"
              @input="onFechaFinCambiada()"
              locale="es-MX"
              placeholder="Selecciona una fecha"
              icon="calendar-today"
              close-on-click
            >
            </b-datepicker>
          </b-field>
        </b-field>
        <b-table
          :data="vehiculos"
          :loading="cargando"
          :mobile-cards="true"
          hoverable
        >
          <b-table-column
            searchable
            field="descripcion"
            label="Descripción"
            v-slot="props"
            sortable
          >
            {{ props.row.descripcion }}
          </b-table-column>
          <b-table-column
            searchable
            field="placas"
            label="Placas"
            sortable
            v-slot="props"
          >
            {{ props.row.placas }}
          </b-table-column>
          <b-table-column
            field="fechaEntrada"
            label="Entrada"
            v-slot="props"
            sortable
          >
            {{ props.row.fechaEntrada | formatearFecha }}
          </b-table-column>
          <b-table-column
            field="fechaSalida"
            label="Salida"
            v-slot="props"
            sortable
          >
            <b-button
              @click="marcarSalida(props.row)"
              v-show="!props.row.fechaSalida"
              type="is-info"
              >Marcar salida</b-button
            >
            <template v-if="props.row.fechaSalida">
              {{ props.row.fechaSalida | formatearFecha }}
              <br />
              <b-icon icon="clock"></b-icon>
              {{ tiempoTranscurrido(props.row) | minutosAHorasYMinutos }}
            </template>
          </b-table-column>
          <b-table-column field="id" label="Imprimir ticket" v-slot="props">
            <b-button
              class="mb-1"
              @click="imprimirTicketEntrada(props.row)"
              type="is-info"
              >Entrada</b-button
            >
            &nbsp;
            <b-button
              v-show="props.row.fechaSalida"
              @click="imprimirTicketSalida(props.row)"
              type="is-primary"
              >Salida</b-button
            >
          </b-table-column>
          <template #empty>
            <div class="has-text-centered">No hay registros</div>
          </template>
        </b-table>
      </div>
    </div>
  </div>
</template>
<script>
import VehiculosService from "../services/VehiculosService";
import TicketService from "../services/TicketService";
import Utiles from "../services/Utiles";
import Cobrar from "./Cobrar";
import DialogosService from "../services/DialogosService";
import PagosService from "../services/PagosService";
import AjustesImpresoraService from "../services/AjustesImpresoraService";
export default {
  components: { Cobrar },
  data: () => ({
    vehiculos: [],
    cargando: false,
    fechaInicio: new Date(),
    fechaFin: new Date(),
    formateadorFecha: Utiles.formatearFechaSegunLocale,
    mostrarModalCobrar: false,
    vehiculoParaCobrar: {},
    impresoraSeleccionada: "",
  }),
  async mounted() {
    await this.obtenerImpresora();
    await this.obtenerVehiculos();
  },
  methods: {
    async obtenerImpresora() {
      this.impresoraSeleccionada = await AjustesImpresoraService.obtenerImpresora();
    },
    async imprimirTicketEntrada(vehiculo) {
      try {
        await TicketService.imprimirTicketEntrada(
          vehiculo,
          this.impresoraSeleccionada
        );
        DialogosService.mostrarNotificacionExito(
          "Ticket impreso correctamente"
        );
      } catch (e) {
        DialogosService.mostrarNotificacionError(
          "Error imprimiendo ticket. Asegúrese de haber configurado su impresora y de que el plugin se esté ejecutando"
        );
      }
    },
    async imprimirTicketSalida(vehiculo) {
      const detallesPago = await PagosService.obtenerPagoPorIdVehiculo(
        vehiculo.id
      );
      const { pago, minutos } = detallesPago;
      try {
        await TicketService.imprimirTicketSalida(
          vehiculo,
          this.impresoraSeleccionada,
          minutos,
          pago
        );
        DialogosService.mostrarNotificacionExito(
          "Ticket impreso correctamente"
        );
      } catch (e) {
        DialogosService.mostrarNotificacionError(
          "Error imprimiendo ticket. Asegúrese de haber configurado su impresora y de que el plugin se esté ejecutando"
        );
      }
    },
    tiempoTranscurrido(vehiculo) {
      const { fechaEntrada, fechaSalida } = vehiculo;
      return Utiles.milisegundosAMinutos(
        Utiles.restarFechasComoCadenas(fechaEntrada, fechaSalida)
      );
    },
    async onCobroTerminado() {
      await this.obtenerVehiculos();
    },
    onFechaInicioCambiada() {
      // Hay que ocultarlo cada que se selecciona una fecha, porque no se oculta automáticamente
      this.$refs.seleccionadorFechaInicio.toggle();
      this.onFechaCambiada();
    },
    onFechaFinCambiada() {
      // Hay que ocultarlo cada que se selecciona una fecha, porque no se oculta automáticamente
      this.$refs.seleccionadorFechaFin.toggle();
      this.onFechaCambiada();
    },
    onFechaCambiada() {
      this.obtenerVehiculos();
    },
    async marcarSalida(vehiculo) {
      this.vehiculoParaCobrar = vehiculo;
      this.mostrarModalCobrar = true;
    },
    async obtenerVehiculos() {
      this.cargando = true;
      const fechaInicio = Utiles.formatearFechaAInicioDeDia(
        this.fechaInicio,
        "T"
      );
      const fechaFin = Utiles.formatearFechaAFinDeDia(this.fechaFin, "T");
      this.vehiculos = await VehiculosService.obtenerVehiculos(
        fechaInicio,
        fechaFin
      );
      this.cargando = false;
    },
  },
};
</script>