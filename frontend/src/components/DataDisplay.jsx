import { useEffect, useState } from "react";
import axios from "axios";

const DataDisplay = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/data")
      .then(response => setData(response.data))
      .catch(error => console.error("Error fetching data:", error));
  }, []);

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="bg-white shadow-lg rounded-lg p-6 max-w-lg w-full">
        <h2 className="text-2xl font-bold mb-4 text-center">Submitted Data</h2>

        {data.length > 0 ? (
          <div className="space-y-2">
            {data.map((item) => (
              <div key={item.id} className="border p-3 rounded-md bg-gray-50">
                <p><strong>Name:</strong> {item.name}</p>
                <p><strong>Email:</strong> {item.email}</p>
              </div>
            ))}

            {/* Download Buttons */}
            <div className="mt-4 flex justify-between">
              <button 
                onClick={() => window.open("http://localhost:8080/download/csv")} 
                className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 transition"
              >
                Download CSV
              </button>
              <button 
                onClick={() => window.open("http://localhost:8080/download/pdf")} 
                className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition"
              >
                Download PDF
              </button>
            </div>
          </div>
        ) : (
          <p className="text-center text-gray-500">No data found...</p>
        )}
      </div>
    </div>
  );
};

export default DataDisplay;
