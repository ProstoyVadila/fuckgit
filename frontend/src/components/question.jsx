
import './question.css';
import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Question = () => {
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
        axios.get("http://localhost:8080/questions")
            .then(res => {
                const payload = res.data.payload;  
                console.log(payload.root);
                setQuestions(payload.root)
            })
            .catch(err => {
                console.log(err)
            });
    }, []);

    return (
        <div className="question">
            <h1>{questions.body}</h1>
            {/* <Link to="/solutions"> */}
            <button className="btn btn-yes" value="yes" onClick={() => {setQuestions(questions.right)}}>
                    ДА
                </button>
            {/* </Link> */}
            <button className="btn btn-no" value="no" onClick={() => setQuestions(questions.left)}>
                Вроде нет
            </button>

            <Link to="/">
                <button className="btn btn-previous" value="prev">
                    Назад
                </button>
            </Link>
        </div>
    )
}

export default Question;
