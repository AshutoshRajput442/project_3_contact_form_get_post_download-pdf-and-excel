// import DataDisplay from "../components/DataDisplay";

// const SubmittedData = () => {
//   return (
//     <div>
//       <h1>Submitted Details</h1>
//       <DataDisplay />
//     </div>
//   );
// };

// export default SubmittedData;
import { useEffect, useState } from "react";
import axios from "axios";

const SubmittedData = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    axios.get("http://localhost:8080/data")
      .then(response => setData(response.data))
      .catch(error => console.error("Error fetching data:", error));
  }, []);

  return (
    <div className="container">
      <h2>Submitted Data</h2>
      {data ? (
        <div>
          <p><strong>Name:</strong> {data.name}</p>
          <p><strong>Email:</strong> {data.email}</p>
          <button onClick={() => window.open("http://localhost:8080/download/csv")}>Download CSV</button>
          <button onClick={() => window.open("http://localhost:8080/download/pdf")}>Download PDF</button>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default SubmittedData;
