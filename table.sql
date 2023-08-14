CREATE TABLE `users` (
  `userId` int(11) AUTO_INCREMENT PRIMARY KEY,
  `email` varchar(225) NOT NULL,
  `phoneNumber` varchar(12) NOT NULL,
  `city` varchar(50) NOT NULL,
  `zip` char(5) NOT NULL,
  `message` text NOT NULL,
  `password` varchar(225) NOT NULL,
  `username` varchar(30) NOT NULL,
  `address` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `admin` (
  `id` int(11) AUTO_INCREMENT PRIMARY KEY,
  `email` varchar(225) NOT NULL,
  `password` varchar(225) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `cars` (
    `carId` int(11) AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(50) NOT NULL,
    `carType` varchar(50) NOT NULL, 
    `rating` ENUM('10%', '20%', '30%', '40%', '50%', '60%', '70%', '80%', '90%', '100%'),
    `fuel` ENUM('PERTAMAX RACING', 'PERTAMAX TURBO', 'PERTAMAX', 'PERTALITE', 'PREMIUM', 'PERTAMINA DEX', 'DEXLITE', 'SOLAR'),
    `image` varchar(225) NOT NULL,
    `hourRate` DOUBLE(20, 2) NULL DEFAULT NULL,
    `dayRate` DOUBLE(20, 2) NULL DEFAULT NULL,
    `monthRate` DOUBLE(20, 2) NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



create TABLE `orders` (
  `orderId` int(11) NOT NULL,
  `pickUpLoc` varchar(255) NOT NULL,
  `dropOffLoc` varchar(255) NOT NULL,
  `pickUpDate` date NOT NULL,
  `dropOffDate` date NOT NULL,
  `pickUpTime` time NOT NULL,
  `carId` int(11) NOT NULL,
  `userId` int(11) NOT NULL,
  `adminId` int(11) NOT NULL,
  FOREIGN KEY (carId) REFERENCES cars (carId),
  FOREIGN KEY (userId) REFERENCES users (userId),
  FOREIGN KEY (adminId) REFERENCES admin (id)
)