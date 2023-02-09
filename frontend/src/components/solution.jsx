import React from "react";
import { Link } from "react-router-dom";

const Solution = () => {
    return (
        <div>
            <h3>Вот что ты ищешь?</h3>
            <p>Решение</p>
            <p>
                <code>git reset --soft ~HEAD</code>
            </p>

            <p>Помогло?</p>
            <Link to="/">
                <button class="btn btn-no" value="no">
                    Да
                </button>
            </Link>
            <button class="btn btn-no" value="no">
                Нет
            </button>

            <Link to="/questions">
                <button class="btn btn-previous" value="prev">
                    Назад
                </button>
            </Link>
        </div>
    );
}

export default Solution;