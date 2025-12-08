/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   estimatePrice.cpp                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: mbatty <mbatty@student.42angouleme.fr>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2025/12/08 13:21:09 by mbatty            #+#    #+#             */
/*   Updated: 2025/12/08 13:21:14 by mbatty           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "estimatePrice.hpp"

double	estimatePrice(double mileage, double theta0, double theta1)
{
	return (theta0 + (theta1 * mileage));
}
