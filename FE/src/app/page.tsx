"use client"

import React, {useEffect, useState} from 'react';

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'


import { Button, Card, ListGroup, ListGroupItem } from 'react-bootstrap';

interface Kols {
	kolID: number;
	userProfileID: number;
	language: string;
	education: string;
	expectedSalary: number;
	expectedSalaryEnable: boolean;
	channelSettingTypeID: number;
	iDFrontURL: string;
	iDBackURL: string;
	portraitURL: string;
	rewardID: number;
	paymentMethodID: number;
	testimonialsID: number;
	verificationStatus: boolean;
	enabled: boolean;
	activeDate: Date;
	active: boolean;
	createdBy: string;
	createdDate: Date;
	modifiedBy: string;
	modifiedDate: Date;
	isRemove: boolean;
	isOnBoarding: boolean;
	code: string;
	portraitRightURL: string;
	portraitLeftURL: string;
	livenessStatus: boolean;
}

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [Kols , setKols] = useState<Kols[]>([]);  

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
    useEffect(() => {
		const fetchData = async() => {
			try {
				const response = await fetch("http://localhost:8081/kols?index=1&size=5")
				const data = await response.json();
				setKols(data.kol);
			}catch(error){
			console.error('Error fetching Kols:', error);
		}
	};
		fetchData();
    }, []);

	console.log(Kols);
    return (
        <>
		<div>
		<h2>KOLs</h2>
		<table className="table table-striped">
          <thead>
            <tr>
              <th>KOL ID</th>
              <th>User Profile ID</th>
              <th>Language</th>
              <th>Education</th>
			  <th scope="col">Expected Salary</th>
			<th scope="col">Expected Salary Enable</th>
			<th scope="col">Channel Setting Type ID</th>
			<th scope="col">ID Front URL</th>
			<th scope="col">ID Back URL</th>
			<th scope="col">Portrait URL</th>
			<th scope="col">Reward ID</th>
			<th scope="col">Payment Method ID</th>
			<th scope="col">Testimonials ID</th>
			<th scope="col">Verification Status</th>
			<th scope="col">Enabled</th>
			<th scope="col">Active Date</th>
			<th scope="col">Active</th>
			<th scope="col">Created By</th>
			<th scope="col">Created Date</th>
			<th scope="col">Modified By</th>
			<th scope="col">Modified Date</th>
			<th scope="col">Is Remove</th>
			<th scope="col">Is On Boarding</th>
			<th scope="col">Code</th>
			<th scope="col">Portrait Right URL</th>
			<th scope="col">Portrait Left URL</th>
			<th scope="col">Liveness Status</th>
            </tr>
          </thead>
          <tbody>
            {Kols.map(kol => (
              <tr key={kol.kolID}>
                <td>{kol.kolID}</td>
                <td>{kol.userProfileID}</td>
                <td>{kol.language}</td>
                <td>{kol.education}</td>
				<td>{kol.expectedSalary}</td>
				<td>{kol.expectedSalaryEnable.toString()}</td>
				<td>{kol.channelSettingTypeID}</td>
				<td>{kol.iDFrontURL}</td>
				<td>{kol.iDBackURL}</td>
				<td>{kol.portraitURL}</td>
				<td>{kol.rewardID}</td>
				<td>{kol.paymentMethodID}</td>
				<td>{kol.testimonialsID}</td>
				<td>{kol.verificationStatus.toString()}</td>
				<td>{kol.enabled.toString()}</td>
				<td>{kol.activeDate.toString()}</td>
				<td>{kol.active.toString()}</td>
				<td>{kol.createdBy}</td>
				<td>{kol.createdDate.toString()}</td>
				<td>{kol.modifiedBy}</td>
				<td>{kol.modifiedDate.toString()}</td>
				<td>{kol.isRemove.toString()}</td>
				<td>{(kol.isOnBoarding).toString()}</td>
				<td>{kol.code}</td>
				<td>{kol.portraitRightURL}</td>
				<td>{kol.portraitLeftURL}</td>
				<td>{kol.livenessStatus.toString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
		</div>
        </>
    )
};

export default Page;