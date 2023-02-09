import React from "react";
import { Link } from "react-router-dom";

const Header = () => {

    function askQuestion() {
        console.log("I'm asking a question!");
    }

    return (
        <div className="Header">
            <h1 class="titile">Fuck Git</h1>
            <p class="title title-description">Мы оба знаем, зачем ты тут. Так чего тянуть?</p>

        <Link to="/questions">
            <button class="btn help-me" onClick={askQuestion}>
                Помоги мне!
            </button>
        </Link>
            
        </div>
    );
}

export default Header;