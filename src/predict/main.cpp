/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.cpp                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mbatty <mbatty@student.42angouleme.fr>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2025/12/08 13:07:40 by mbatty            #+#    #+#             */
/*   Updated: 2025/12/08 15:56:47 by mbatty           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <iostream>
#include <string>
#include <sys/signal.h>
#include "estimatePrice.hpp"
#include "color.hpp"

int	main(int ac, char **av)
{
	if (ac != 3)
	{
		std::cerr << "Usage: ./Predict theta0 theta1" << std::endl;
		return (1);
	}

	double	theta0;
	double	theta1;
	try {
		theta0 = std::stod(av[1]);
		theta1 = std::stod(av[2]);
	} catch (const std::exception &e) {
		std::cerr << COLOR_RED "Error failed to parse thetas: " << e.what() << COLOR_RESET << std::endl;
		return (1);
	}

	std::string	str;
	double		mileage;
	while (42)
	{
		std::cout << "Mileage: " << std::flush;
		if (!std::getline(std::cin, str))
			break ;
		try {
			mileage = std::stod(str);			
		} catch (const std::exception &e) {
			std::cerr << COLOR_GRAY "Error: " << e.what() << COLOR_RESET << std::endl;
			continue ;
		}
		std::cout << estimatePrice(mileage, theta0, theta1) << std::endl;
	}
	return (0);
}
