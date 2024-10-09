import moveIcon from '../homepage_components../src/move_icon.png';
import PropTypes from 'prop-types'

function Card(props) {
    return (
        <div className="card">
            <h4>{props.title}</h4>
            <p className="category">
                {props.text}
            </p>
            <div className="card-group">
                <div className="group-div">
                    {props.category}
                </div>
                <div className="card-move">
                    <button className="move-button">
                        <img src={moveIcon} alt="move icon" height="25px"></img>
                    </button>
                </div>
            </div>
        </div>
    );
}

Card.propTypes = {
    title: PropTypes.string,
    text: PropTypes.string,
    category: PropTypes.string
}

Card.defaultProps = {
    title: "Time tracking",
    text: "Work in department",
    category: "TSD-15"
}

export default Card