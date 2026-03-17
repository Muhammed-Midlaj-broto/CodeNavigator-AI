import {
Chart as ChartJS,
CategoryScale,
LinearScale,
BarElement,
Tooltip,
Legend
} from "chart.js";

import { Bar } from "react-chartjs-2";

ChartJS.register(
CategoryScale,
LinearScale,
BarElement,
Tooltip,
Legend
);

export default function LanguageChart({ languages }) {

if(!languages) return null;

const labels = Object.keys(languages);
const values = Object.values(languages);

const data = {
labels: labels,
datasets: [
{
label: "Files",
data: values,
backgroundColor: "#38bdf8",
borderRadius: 5
}
]
};

const options = {
indexAxis: "y",
plugins:{
legend:{display:false}
}
};

return <Bar data={data} options={options} />;
}