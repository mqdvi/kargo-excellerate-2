import React, { Component } from 'react';
import { Link } from 'react-router-dom';

export default function Home() {
  return (
    <>
      <div className="min-h-full">
        <main>
          <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
            <div className="px-4 py-6 sm:px-0">
                <Link to={"/"}>
                    <p className='h3 px-10 py-2'>1. Main Page</p>
                </Link>
                <Link to={"/get-list-berita"}>
                    <p className='h3 px-10 py-2'>2. Get List Berita</p>
                </Link>
            </div>
          </div>
        </main>
      </div>
    </>
  )
}
