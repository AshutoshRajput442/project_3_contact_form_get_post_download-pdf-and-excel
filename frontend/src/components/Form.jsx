// import { useState } from "react";
// import { useNavigate } from "react-router-dom";
// import axios from "axios";

// const Form = () => {
//   const [formData, setFormData] = useState({ name: "", email: "" });
// //   const navigate = useNavigate();

//   const handleChange = (e) => {
//     setFormData({ ...formData, [e.target.name]: e.target.value });
//   };

//   const handleSubmit = async (e) => {
//     e.preventDefault();
//     try {
//       await axios.post("http://localhost:8080/submit", formData);
//       alert("Data submitted successfully! ✅");  // Success alert
//       navigate("/submitted");
//     } catch (error) {
//       alert("Error submitting data ❌");
//       console.error(error);
//     }
//   };

//   return (
//     <div className="flex justify-center items-center min-h-screen bg-gray-100">
//       <div className="bg-white shadow-lg rounded-lg p-8 max-w-md w-full">
//         <h2 className="text-2xl font-bold mb-4 text-center">User Form</h2>
//         <form onSubmit={handleSubmit} className="space-y-4">
//           <input
//             type="text"
//             name="name"
//             placeholder="Enter your name"
//             value={formData.name}
//             onChange={handleChange}
//             className="w-full p-2 border rounded-md"
//             required
//           />
//           <input
//             type="email"
//             name="email"
//             placeholder="Enter your email"
//             value={formData.email}
//             onChange={handleChange}
//             className="w-full p-2 border rounded-md"
//             required
//           />
//           <button
//             type="submit"
//             className="w-full bg-blue-500 hover:bg-blue-600 text-white py-2 rounded-md transition"
//           >
//             Submit
//           </button>
//         </form>
//       </div>
//     </div>
//   );
// };

// export default Form;
import { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom"; // 👈 Import Navigation

const Form = () => {
  const [formData, setFormData] = useState({ name: "", email: "" });
  const navigate = useNavigate(); // 👈 Initialize Navigation

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post("http://localhost:8080/submit", formData);
      alert("Data submitted successfully! ✅");
      navigate("/submitted"); // 👈 Redirect to SubmittedData page
    } catch (error) {
      console.error("Error submitting form:", error);
      alert("Submission failed! ❌");
    }
  };

  return (
    <div className="container">
      <h2>Submit Your Details</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" name="name" placeholder="Enter Name" onChange={handleChange} required />
        <input type="email" name="email" placeholder="Enter Email" onChange={handleChange} required />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default Form;
