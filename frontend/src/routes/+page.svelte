<script lang="ts">
	import type { PageData } from './$types';
	import toast, { Toaster } from 'svelte-french-toast';
	export let data: PageData;

	let selectedSeats: { row: number; seat: number }[] = [];
	let totalPrice: number = 0;
	let seatCount: number = 0;

	let booker_name: string = '';
	let booker_phone: string = '';

	let movie = { name: data.movies[0].name, price: data.movies[0].amount, index: 0 }; // Default movie
	$: total = totalPrice;
	$: count = seatCount;

	let seats = data.movies[0].seats;

	$: allOccupied = seats.flatMap((row) => row).every((seat) => seat === 'occupied');

	async function post_boker() {
		if (seatCount === 0) {
			toast.error('Must select seats');
			return;
		}

		if (booker_name === '' || booker_phone === '') {
			toast.error('Must enter name & phone number');
			return;
		}

		if (booker_phone.length !== 10) {
			toast.error('Must enter valid number');
			return;
		}

		// Check if was a voda number
		// let vd: string = booker_phone.slice(1, 3);
		// if (vd !== '74' && vd !== '75' && vd !== '76') {
		// 	toast.error('Must enter voda number only');
		// 	return;
		// }

		const res = await fetch('/api/movie', {
			method: 'POST',
			body: JSON.stringify({
				name: booker_name,
				phone: booker_phone,
				seat: selectedSeats,
				movie: movie.name,
				amount: movie.price,
				seats: seats
			}),
			headers: {
				'content-type': 'application/json'
			}
		});

		const respSeats = await res.json();
		seats = [...respSeats.seats];
		data.movies[movie.index].seats = [...respSeats.seats];
		booker_name = '';
		booker_phone = '';
		seatCount = 0;
		totalPrice = 0;
		selectedSeats = [];
		toast.success('Ticket Booked Successfully');
	}

	function toggleSeat(row: number, seat: number): void {
		if (seats[row][seat] === 'selected') {
			// Deselect seat
			seats[row][seat] = '';
			seatCount -= 1;
			totalPrice -= movie.price;
			selectedSeats = selectedSeats.filter(
				(selectedSeat) => !(selectedSeat.row === row && selectedSeat.seat === seat)
			);
		} else if (seats[row][seat] === '') {
			// Select seat
			seats[row][seat] = 'selected';
			seatCount += 1;
			totalPrice += movie.price;
			selectedSeats.push({ row, seat });
		}
	}

	function handleChange(event: Event): void {
		const target = event.target as HTMLSelectElement;
		movie.name = target.value;

		const selectedMovie = data.movies.find((m) => m.name === movie.name);
		movie.index = data.movies.findIndex((m) => m.name === movie.name);

		seats = [...(selectedMovie?.seats ?? [])];

		movie.price = selectedMovie?.amount ?? 0;

		totalPrice = 0;
		seatCount = 0;
		selectedSeats = [];
		seats.forEach((row, rowIndex) => {
			row.forEach((seat, seatIndex) => {
				if (seat === 'selected') {
					seatCount += 1;
					totalPrice += movie.price;
					selectedSeats.push({ row: rowIndex, seat: seatIndex });
				}
			});
		});
	}
</script>

<body>
	<div class="cinema">
		<Toaster />
		<div class="movie-container">
			<label for="movie">Pick a movie:</label>
			<select id="movie" on:change={handleChange}>
				{#each data.movies as movie}
					<option value={movie.name}>{movie.name} (${movie.amount})</option>
				{/each}
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

		{#if allOccupied === false}
			<p class="text">
				You have selected <span id="count">{count}</span> seats for a price of $<span id="total"
					>{total}</span
				>
			</p>
		{/if}
	</div>

	<div class="classyform">
		{#if allOccupied}
			<div class="sold-out">sold out</div>
		{:else}
			<h3>Enter your details</h3>
			<label>
				Name
				<input bind:value={booker_name} type="text" placeholder="Lugano Abel" />
			</label>
			<label>
				Phone
				<input bind:value={booker_phone} type="text" placeholder="0754123456" />
			</label>
			<button on:click={post_boker}>Buy</button>
		{/if}
	</div>
</body>

<style>
	@import '../styles.css';
</style>
