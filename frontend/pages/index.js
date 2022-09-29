import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import 'tailwindcss/tailwind.css'


export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []); 

  console.log(data?.data)
  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    } 
    catch (error) {
      setError(error);
    }
  }

  const deleteItem = async(id) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/delete/${id}`, {
        method:'DELETE',
      });
      const data = await res.json();
      getData()
    } 
    catch (error) {
      setError(error);
    }
  }

  const updateData = async(id) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/update/${id}`, {
        method:'PUT',
      });
      const data = await res.json();
      getData()
    } 
    catch (error) {
      setError(error);
    }
  }

  return (
    <div class="grid h-screen place-items-center">
      <div>
      {error && <div>Failed to load {error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0 && <p>data kosong</p>
          )
      }
      <Input onSuccess={getData}/>
      {data?.data && data?.data?.map((item, index) => (
        <div key={index}>
          <span class="mr-2">ID: {item.ID} task: {item.task}</span>
          <input type="checkbox" defaultChecked={item.done} class="mr-4"/>
          <button onClick={() => updateData(item.ID)} class="bg-emerald-500 text-white active:bg-emerald-600 font-bold uppercase text-xs px-4 py-2 rounded shadow hover:shadow-md outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button">
            Selesai
          </button>
          <button onClick={() => deleteItem(item.ID)} class="bg-red-500 text-white active:bg-red-600 font-bold uppercase text-xs px-4 py-2 rounded shadow hover:shadow-md outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150" type="button">
            <svg className="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" /></svg>
          </button>
          </div>  
      ))}
      </div>
    </div>  
  )
}


function Input({onSuccess}) {
  const [data, setData] =  useState(null);
  const [error, setError] =  useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      task: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();
      setData(data.message);
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }
  return (
    <div>
      <h1 class="font-mono text-xl">Selamat datang</h1>
      <p class="font-mono">Silahkan input data</p>
      <br />
      <form class="w-full max-w-sm border-2 rounded-md border-white-500" onSubmit={handleSubmit} id="form1">
        <div class="flex items-center border-b border-teal-500 py-2">
          <input class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="Jane Doe" aria-label="text" name='data' />
          <button class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded" type="submit">
          Submit
        </button>
        <button class="flex-shrink-0 border-transparent border-4 text-teal-500 hover:text-teal-800 text-sm py-1 px-2 rounded" type="button">
          Cancel
        </button>
      </div>
      </form>
      {error && <p>error: {error.toString()}</p>}
      {data && <p>success: {data}</p>}
      <hr />
      <br />
    </div>
  )
}

