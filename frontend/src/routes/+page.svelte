<script lang="ts">
  let selectedSeats: { row: number; seat: number }[] = [];
  let totalPrice: number = 0;
  let seatCount: number = 0;

  const seatPrice: { [key: number]: number } = {
    10: 10,
    12: 12,
    8: 8,
    9: 9,
  };

  let movie = { value: 10 }; // Default movie
  $: total = totalPrice;
  $: count = seatCount;

  let seats: (null | "selected" | "occupied")[][] = [
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
    [null, null, null, null, null, null, null, null],
  ];

  function toggleSeat(row: number, seat: number): void {
    if (seats[row][seat] === "selected") {
      // Deselect seat
      seats[row][seat] = null;
      seatCount -= 1;
      totalPrice -= seatPrice[movie.value];
      selectedSeats = selectedSeats.filter(
        (selectedSeat) => !(selectedSeat.row === row && selectedSeat.seat === seat)
      );
    } else if (seats[row][seat] === null) {
      // Select seat
      seats[row][seat] = "selected";
      seatCount += 1;
      totalPrice += seatPrice[movie.value];
      selectedSeats.push({ row, seat });
    }
    seats = [...seats]; 

    console.log(seats)
    console.log(row, seat)
    console.log("selected",selectedSeats)
  }

  function handleChange(event: Event): void {
    const target = event.target as HTMLSelectElement;
    movie.value = parseInt(target.value, 10);
    totalPrice = 0;
    seatCount = 0;
    selectedSeats = [];
    seats.forEach((row, rowIndex) => {
      row.forEach((seat, seatIndex) => {
        if (seat === "selected") {
          seatCount += 1;
          totalPrice += seatPrice[movie.value];
          selectedSeats.push({ row: rowIndex, seat: seatIndex });
        }
      });
    });
  }
</script>
  
  <body>
    <div class="movie-container">
    <label for="movie">Pick a movie:</label>
      <select id="movie" on:change={handleChange}>
        <option value="10">Avengers: Endgame ($10)</option>
        <option value="12">Joker ($12)</option>
        <option value="8">Toy Story 4 ($8)</option>
        <option value="9">The Lion King ($9)</option>
      </select>
    </div>

    <ul class="showcase">
      <li>
        <div class="seat"></div>
        <small>Unselected</small>
      </li>
      <li>
        <div class="seat selected"></div>
        <small>Selected</small>
      </li>
      <li>
        <div class="seat occupied"></div>
        <small>Occupied</small>
      </li>
    </ul>

    <div class="container">
      <div class="screen"></div>
      {#each seats as row, rowIndex}
      <div class="row">
        {#each row as seat, seatIndex}
        <button 
            class="seat {seat}" 
            on:click={() => toggleSeat(rowIndex, seatIndex)}
            title={`Row: ${rowIndex + 1}, Seat: ${seatIndex + 1}`}
        >
            <!-- <div style="font-size: 2;">12</div> -->
        </button>
        {/each}
      </div>
    {/each}
  </div>

      

    <p class="text">
      You have selected <span id="count">{count}</span> seats for a price of $<span
        id="total"
        >{total}</span
      >
    </p>

</body>

<style>
    @import url("https://fonts.googleapis.com/css2?family=Lato&display=swap");

* {
  box-sizing: border-box;
}

body {
  background-color: #242333;
  color: #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 80vh;
  font-family: "Lato", sans-serif;
  margin: 0;
}

.movie-container {
  margin: 40px 0;
}

.movie-container select {
  background-color: #fff;
  border: 0;
  border-radius: 5px;
  font-size: 14px;
  margin-left: 10px;
  padding: 5px 15px 5px 15px;
  -moz-appearance: none;
  -webkit-appearance: none;
  appearance: none;
}

.container {
  perspective: 1000px;
  margin-bottom: 30px;
}

.seat {
  background-color: #444451;
  height: 24px;
  width: 20px;
  margin: 3px;
  border-top-left-radius: 6px;
  border-top-right-radius: 6px;
}

.seat.selected {
  background-color: #6feaf6;
}

.seat.occupied {
  background-color: #fff;
}

.seat:nth-of-type(2) {
  margin-right: 18px;
}

.seat:nth-last-of-type(2) {
  margin-left: 18px;
}

.seat:not(.occupied):hover {
  cursor: pointer;
  transform: scale(1.2);
}

.showcase .seat:not(.occupied):hover {
  cursor: default;
  transform: scale(2);
}

.showcase {
  background: rgba(0, 0, 0, 0.1);
  padding: 5px 10px;
  border-radius: 5px;
  color: #777;
  list-style-type: none;
  display: flex;
  justify-content: space-between;
}

.showcase li {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 10px;
}

.showcase li small {
  margin-left: 2px;
}

.row {
  display: flex;
}

.screen {
  background-color: #fff;
  height: 70px;
  width: 100%;
  margin: 15px 0;
  transform: rotateX(-45deg);
  box-shadow: 0 3px 10px rgba(255, 255, 255, 0.7);
}

p.text {
  margin: 5px 0;
}

p.text span {
  color: #6feaf6;
}


</style>