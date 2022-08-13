import React, { Component } from 'react';
import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import { MenuIcon, XIcon } from '@heroicons/react/outline'
import { ChevronDownIcon } from '@heroicons/react/solid'

function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export class Navbar extends Component {
    myColor(position) {
        if (window.location.pathname === position) {
            return "#FF8300";
        }
    }

    render() {
        return (
            <div className='font-roboto text-navbar text-xl sticky top-0 bg-white drop-shadow-xl z-10'>
                <Popover>
                    <div className="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
                        <div className="flex items-center justify-between h-20 md:justify-start md:space-x-10">

                            <div className="-mr-2 -my-2 md:hidden">
                                <Popover.Button className="bg-white rounded-md p-2 inline-flex items-center justify-center hover:text-oren3 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset">
                                    <span className="sr-only">Open menu</span>
                                    <MenuIcon className="h-6 w-6" aria-hidden="true" />
                                </Popover.Button>
                            </div>

                            <Popover.Group as="nav" className="flex justify-end hidden md:flex space-x-10 ">
                                <a href="/" className="hover:text-oren3" style={{color: this.myColor("/")}}>LMS</a>
                                <a href="/get-list-berita" className="hover:text-oren3" style={{color: this.myColor("/get-list-berita")}}>Shipments</a>
                                <a href="/Transporter/trucks" className="hover:text-oren3" style={{color: this.myColor("/Transporter/trucks")}}>Trucks</a>
                                <a href="/read-artikel" className="hover:text-oren3" style={{color: this.myColor("/read-artikel")}}>Drivers</a>
                            </Popover.Group>
                            
                            <div className="hidden md:flex items-center justify-end">
                                <a href="/register" className="p-1 px-5 rounded-full text-lg text-white bg-oren3 hover:bg-yellow-700">Login/Register</a>
                            </div>
                        </div>
                    </div>

                    <Transition as={Fragment} enter="duration-200 ease-out" enterFrom="opacity-0 scale-95" enterTo="opacity-100 scale-100" leave="duration-100 ease-in" leaveFrom="opacity-100 scale-100" leaveTo="opacity-0 scale-95" >
                        <Popover.Panel focus className="absolute top-0 inset-x-0 p-2 transition transform origin-top-right md:hidden">
                            <div className="rounded-lg shadow-lg ring-1 ring-black ring-opacity-5 bg-white divide-y-2 divide-gray-50">
                                <div className="pt-5 pb-6 px-5">
                                    <div className="flex items-center justify-between">
                                        <div className="-mr-2">
                                            <Popover.Button className="bg-white rounded-md p-2 inline-flex items-center justify-center hover:text-oren3 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500">
                                                <span className="sr-only">Close menu</span>
                                                <XIcon className="h-6 w-6" aria-hidden="true" />
                                            </Popover.Button>
                                        </div>
                                    </div>
                                    <div className="mt-6">
                                        <nav className="grid gap-y-8">
                                            <a href="/" className="-m-3 p-3 flex items-center rounded-md hover:bg-gray-50" >
                                                <span className="ml-3 text-base hover:text-oren3" style={{color: this.myColor("/")}}>LMS</span>
                                            </a>
                                            <a href="/get-list-berita" className="-m-3 p-3 flex items-center rounded-md hover:bg-gray-50" >
                                                <span className="ml-3 text-base hover:text-oren3" style={{color: this.myColor("/get-list-berita")}}>Shipments</span>
                                            </a>
                                            <a href="/Transporter/trucks" className="-m-3 p-3 flex items-center rounded-md hover:bg-gray-50" >
                                                <span className="ml-3 text-base hover:text-oren3" style={{color: this.myColor("/Transporter/trucks")}}>Trucks</span>
                                            </a>
                                            <a href="/read-artikel" className="-m-3 p-3 flex items-center rounded-md hover:bg-gray-50" >
                                                <span className="ml-3 text-base hover:text-oren3" style={{color: this.myColor("/read-artikel")}}>Drivers</span>
                                            </a>
                                        </nav>
                                    </div>
                                </div>
                            </div>
                        </Popover.Panel>
                    </Transition>
                </Popover>
            </div>
        );  
    }
}
export default Navbar;