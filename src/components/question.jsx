
import './question.css';
import React from 'react';
import { Link } from 'react-router-dom';

const Question = () => {

    function decide(event) {
        console.log(event.target.value)
        // alert(event.target.value)
    }

    return (
        <div className="question">
            <h1>Question</h1>
            <p>Are you crazy, man?</p>
            <Link to="/solutions">
                <button class="btn btn-yes" value="yes" onClick={decide}>
                    ДА
                </button>
            </Link>
            <button class="btn btn-no" value="no" onClick={decide}>
                Вроде нет
            </button>

            <Link to="/">
                <button class="btn btn-previous" value="prev" onClick={decide}>
                    Назад
                </button>
            </Link>
        </div>
    )
}

export default Question;
