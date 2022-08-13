import React, { Component } from 'react';
import { Link } from 'react-router-dom';

class Home extends Component {
  
  render() {    
    return (
      <>
        <div className="min-h-full">
          <main>
            <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
              <div className="px-4 py-6 sm:px-0">
                  <Link to={"/login"}>
                      <input 
                          type = 'submit'
                          value = 'Login'
                          className='bg-yellow-100 px-5 py-2'
                      />
                  </Link>
              </div>
            </div>
          </main>
        </div>
      </>
    )
  }
}

export default Home